package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateProjectGroup(spaceName string, projectGroup *model.ProjectGroup) {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	projectGroupService := octopus.NewProjectGroupService(octopusHost, octopusApiKey)

	space := spaceService.GetByName(spaceName)
	if space == nil {
		panic("Space Not Found")
	}
	spaceId := space.Id

	projectGroup.SetSpaceId(spaceId)
	pg := projectGroupService.GetByName(spaceId, projectGroup.Name)

	if pg == nil {
		projectGroupService.Create(projectGroup)
	} else {
		projectGroupId := pg.Id
		projectGroup.SetId(projectGroupId)
		projectGroupService.Update(projectGroupId, projectGroup)
	}
}
