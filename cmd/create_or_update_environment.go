package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateEnvironment(spaceName string, environment *model.Environment) {

	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	service := octopus.NewEnvironmentService(octopusHost, octopusApiKey)

	s := spaceService.GetByName(spaceName)
	if s == nil {
		panic("Space does not exist")
	}
	createOrUpdateEnvironment(s.Id, environment, service)
}

func createOrUpdateEnvironment(spaceId string, environment *model.Environment, service *octopus.EnvironmentService) string {
	environment.SetSpaceId(spaceId)
	e := service.GetByName(spaceId, environment.Name)
	if e == nil {
		e = service.Create(environment)
	} else {
		environment.SetId(e.Id)
		e = service.Update(e.Id, environment)
	}
	return e.Id
}
