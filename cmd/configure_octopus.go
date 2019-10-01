package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func ConfigureOctopus(space *model.Space, octopusEnvironment []model.OctopusEnvironment, envNames, phasesOrder []string, automaticTargets, optionalTargets map[string][]string, lifecycle *model.Lifecycle, feed *model.Feed, actionTemplateInputs []model.ActionTemplateInput) {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	accountService := octopus.NewAccountService(octopusHost, octopusApiKey)
	actionTemplateService := octopus.NewActionTemplateService(octopusHost, octopusApiKey)
	environmentService := octopus.NewEnvironmentService(octopusHost, octopusApiKey)
	feedService := octopus.NewFeedService(octopusHost, octopusApiKey)
	lifecycleService := octopus.NewLifecycleService(octopusHost, octopusApiKey)
	machineService := octopus.NewMachineService(octopusHost, octopusApiKey)

	spaceId := createOrUpdateSpace(space, spaceService)

	for _, oe := range octopusEnvironment {
		envId := createOrUpdateEnvironment(spaceId, oe.Environment, environmentService)
		accountId := createOrUpdateAccount(spaceId, []string{envId}, oe.Account, accountService)
		createOrUpdateMachine(spaceId, []string{envId}, accountId, oe.Machine, machineService)
	}

	createOrUpdateLifecycle(spaceId, phasesOrder, envNames, automaticTargets, optionalTargets, lifecycle, environmentService, lifecycleService)
	createOrUpdateFeed(spaceId, feed, feedService)

	for _, ati := range actionTemplateInputs {
		createOrUpdateActionTemplate(spaceId, ati.Name, ati.Desc, ati.Type, feedService, actionTemplateService)
	}
}
