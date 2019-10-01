package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateAccount(spaceName string, envNames []string, account *model.Account) {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	envService := octopus.NewEnvironmentService(octopusHost, octopusApiKey)
	accountService := octopus.NewAccountService(octopusHost, octopusApiKey)

	s := spaceService.GetByName(spaceName)
	if s == nil {
		panic("Space does not exist")
	}

	envIds := make([]string, len(envNames))
	for i, envName := range envNames {
		e := envService.GetByName(s.Id, envName)

		if e == nil {
			panic("Environment doesn't exist")
		}
		envIds[i] = e.Id
	}
	createOrUpdateAccount(s.Id, envIds, account, accountService)
}

func createOrUpdateAccount(spaceId string, envIds []string, account *model.Account, accountService *octopus.AccountService) string {
	account.SetSpaceId(spaceId)
	account.SetEnvironmentIds(envIds)

	acc := accountService.GetByName(spaceId, account.Name)

	if acc == nil {
		acc = accountService.Create(account)
	} else {
		account.SetId(acc.Id)
		acc = accountService.Update(acc.Id, account)
	}

	return acc.Id
}
