package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateSpace(space *model.Space) {

	service := octopus.NewSpaceService(octopusHost, octopusApiKey)

	createOrUpdateSpace(space, service)

}

func createOrUpdateSpace(space *model.Space, spaceService *octopus.SpaceService) string {
	s := spaceService.GetByName(space.Name)
	if s == nil {
		s = spaceService.Create(space)
	} else {
		space.SetId(s.Id)
		space.UpdateSpaceManaget(s.SpaceManagersTeams)
		s = spaceService.Update(space.Id, space)
	}
	return s.Id
}
