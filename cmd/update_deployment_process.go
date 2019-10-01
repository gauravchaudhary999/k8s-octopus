package cmd

import (
	"strconv"

	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func UpdateDeploymentProcess(spaceName string, projectName string, templateNames []string) {

	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	projectService := octopus.NewProjectService(octopusHost, octopusApiKey)
	actionTemplateService := octopus.NewActionTemplateService(octopusHost, octopusApiKey)
	deploymentProcessService := octopus.NewDeploymentProcessService(octopusHost, octopusApiKey)

	space := spaceService.GetByName(spaceName)
	if space == nil {
		panic("Space Not Found")
	}
	spaceId := space.Id

	project := projectService.GetByName(spaceId, projectName)

	if project == nil {
		panic("Project Not Found")
	}
	deploymentProcess := deploymentProcessService.Get(spaceId, project.DeploymentProcessId)
	steps := make([]model.Step, len(templateNames))

	for index, templateName := range templateNames {
		steps[index] = makeStep(spaceId, templateName, steps, actionTemplateService)
	}

	dp := model.DeploymentProcess{
		Id:        deploymentProcess.Id,
		SpaceId:   spaceId,
		ProjectId: project.Id,
		Version:   deploymentProcess.Version,
		Steps:     steps,
	}

	deploymentProcessService.Update(deploymentProcess.Id, &dp)

}

func isStepAlreadyPresent(steps []model.Step, name string) bool {
	for _, s := range steps {
		if s.Name == name {
			return true
		}
	}
	return false
}

func makeStep(spaceId string, templateName string, steps []model.Step, actionTemplateService *octopus.ActionTemplateService) model.Step {
	actionTemplate := actionTemplateService.GetByName(spaceId, templateName)

	if actionTemplate == nil {
		panic("Action Template not found")
	}

	properties := actionTemplate.Properties
	properties["Octopus.Action.Template.Id"] = actionTemplate.Id
	properties["Octopus.Action.Template.Version"] = strconv.Itoa(actionTemplate.Version)

	action := model.Action{
		Name:                          actionTemplate.Name,
		ActionType:                    actionTemplate.ActionType,
		Properties:                    properties,
		Packages:                      actionTemplate.Packages,
		IsDisabled:                    false,
		CanBeUsedForProjectVersioning: true,
		IsRequired:                    false,
	}
	props := make(map[string]string)
	props["Octopus.Action.TargetRoles"] = "admin"
	props["Octopus.Action.MaxParallelism"] = "1"
	step := model.Step{
		Name:               actionTemplate.Name,
		Actions:            []model.Action{action},
		PackageRequirement: "LetOctopusDecide",
		Condition:          "Success",
		StartTrigger:       "StartAfterPrevious",
		Properties:         props,
	}
	return step
}
