package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateProject(spaceName string, projectGroupName string, lifecycleName string, project *model.Project) {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	projectService := octopus.NewProjectService(octopusHost, octopusApiKey)
	projectGroupService := octopus.NewProjectGroupService(octopusHost, octopusApiKey)
	lifecycleService := octopus.NewLifecycleService(octopusHost, octopusApiKey)

	space := spaceService.GetByName(spaceName)
	if space == nil {
		panic("Space Not Found")
	}
	spaceId := space.Id

	projectGroup := projectGroupService.GetByName(spaceId, projectGroupName)
	if projectGroup == nil {
		panic("Project group not found")
	}

	projectGroupId := projectGroup.Id

	lifecycle := lifecycleService.GetByName(spaceId, lifecycleName)
	if lifecycle == nil {
		panic("Lifecycle not found")
	}

	lifecycleId := lifecycle.Id

	project.SetSpaceId(spaceId)
	project.SetProjectGroupId(projectGroupId)
	project.SetLifecycleId(lifecycleId)

	p := projectService.GetByName(spaceId, project.Name)
	if p == nil {
		projectService.Create(project)
	} else {
		projectId := p.Id
		project.SetId(projectId)
		projectService.Update(projectId, project)
	}

}
