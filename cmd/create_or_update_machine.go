package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateMachine(spaceName string, envNames []string, accountName string, machine *model.Machine) {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	envService := octopus.NewEnvironmentService(octopusHost, octopusApiKey)
	accountService := octopus.NewAccountService(octopusHost, octopusApiKey)
	machineService := octopus.NewMachineService(octopusHost, octopusApiKey)

	space := spaceService.GetByName(spaceName)
	if space == nil {
		panic("Space Not Found")
	}

	envIds := make([]string, len(envNames))
	for i, envName := range envNames {
		e := envService.GetByName(space.Id, envName)

		if e == nil {
			panic("Environment doesn't exist")
		}
		envIds[i] = e.Id
	}

	account := accountService.GetByName(space.Id, accountName)

	if account == nil {
		panic("Account does not exist")
	}
	createOrUpdateMachine(space.Id, envIds, account.Id, machine, machineService)
}

func createOrUpdateMachine(spaceId string, envIds []string, accountId string, machine *model.Machine, machineService *octopus.MachineService) string {
	machine.SetSpaceId(spaceId)
	machine.SetEnvironmentIds(envIds)
	machine.Endpoint.Authentication.SetAccountId(accountId)
	m := machineService.GetByName(spaceId, machine.Name)

	if m == nil {
		m = machineService.Create(machine)
	} else {
		machine.SetId(m.Id)
		m = machineService.Update(m.Id, machine)
	}

	return m.Id
}
