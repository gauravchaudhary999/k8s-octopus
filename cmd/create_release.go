package cmd

import (
	"time"

	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateRelease(spaceName, projectName, serviceChartVersion, helmClientVersion, servicePackageVersion, apiGatewayToolVersion string) {

	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	projectService := octopus.NewProjectService(octopusHost, octopusApiKey)
	deploymentProcessService := octopus.NewDeploymentProcessService(octopusHost, octopusApiKey)
	releaseService := octopus.NewReleaseService(octopusHost, octopusApiKey)

	space := spaceService.GetByName(spaceName)
	if space == nil {
		panic("Space Not Found")
	}
	spaceId := space.Id

	project := projectService.GetByName(spaceId, projectName)

	if project == nil {
		panic("Project Not Found")
	}
	projectId := project.Id

	packagesVersionMapping := make(map[string]string)

	ifPresentThenAdd(packagesVersionMapping, "ServiceChart", serviceChartVersion)
	ifPresentThenAdd(packagesVersionMapping, "HelmClient", helmClientVersion)
	ifPresentThenAdd(packagesVersionMapping, "ServicePackage", servicePackageVersion)
	ifPresentThenAdd(packagesVersionMapping, "ApiGatewayTool", apiGatewayToolVersion)

	deploymentProcess := deploymentProcessService.Get(spaceId, project.DeploymentProcessId)

	var selectedPackages []model.SelectedPackage
	for _, step := range deploymentProcess.Steps {
		for _, action := range step.Actions {

			for _, aPackage := range action.Packages {
				sp := model.SelectedPackage{
					StepName:             step.Name,
					ActionName:           action.Name,
					Version:              packagesVersionMapping[aPackage.Name],
					PackageReferenceName: aPackage.Name,
				}
				selectedPackages = append(selectedPackages, sp)
			}

		}
	}

	release := model.Release{
		SpaceId:          spaceId,
		ProjectId:        projectId,
		Assembled:        time.Now().String(),
		Version:          servicePackageVersion,
		SelectedPackages: selectedPackages,
	}

	releaseService.Create(&release)
}

func ifPresentThenAdd(packagesVersionMapping map[string]string, key, value string) {

	if value != "" {
		packagesVersionMapping[key] = value
	}
}
