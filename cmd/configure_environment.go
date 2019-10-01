package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func ConfigureEnvironment(spaceName string, env *model.Environment, acc *model.Account, machine *model.Machine) string {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	environmentService := octopus.NewEnvironmentService(octopusHost, octopusApiKey)
	accountService := octopus.NewAccountService(octopusHost, octopusApiKey)
	machineService := octopus.NewMachineService(octopusHost, octopusApiKey)

	s := spaceService.GetByName(spaceName)
	if s == nil {
		panic("Space does not exist")
	}
	spaceId := s.Id

	envId := createOrUpdateEnvironment(spaceId, env, environmentService)
	accountId := createOrUpdateAccount(spaceId, []string{envId}, acc, accountService)
	createOrUpdateMachine(spaceId, []string{envId}, accountId, machine, machineService)
	return envId
}
