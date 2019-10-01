package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateLifecycle(spaceName string, phasesOrder, envNames []string, automaticTargets, optionalTargets map[string][]string, lifecycle *model.Lifecycle) {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	envService := octopus.NewEnvironmentService(octopusHost, octopusApiKey)
	lifecycleService := octopus.NewLifecycleService(octopusHost, octopusApiKey)

	s := spaceService.GetByName(spaceName)
	if s == nil {
		panic("Space does not exist")
	}
	createOrUpdateLifecycle(s.Id, phasesOrder, envNames, automaticTargets, optionalTargets, lifecycle, envService, lifecycleService)
}

func createOrUpdateLifecycle(spaceId string, phasesOrder, envNames []string, automaticTargets, optionalTargets map[string][]string, lifecycle *model.Lifecycle, envService *octopus.EnvironmentService, lifecycleService *octopus.LifecycleService) {

	envIds := make(map[string]string)
	for _, envName := range envNames {
		e := envService.GetByName(spaceId, envName)
		if e == nil {
			panic("Environment doesn't exist")
		}
		envIds[e.Name] = e.Id
	}

	lifecycle.SetSpaceId(spaceId)
	phases := make([]model.Phase, len(phasesOrder))
	for index, phaseName := range phasesOrder {
		aEnvNames := automaticTargets[phaseName]
		oEnvNames := optionalTargets[phaseName]
		phases[index] = model.Phase{Name: phaseName, AutomaticDeploymentTargets: transformNamesToIds(envIds, aEnvNames), OptionalDeploymentTargets: transformNamesToIds(envIds, oEnvNames)}
	}
	lifecycle.SetPhases(&phases)
	l := lifecycleService.GetByName(spaceId, lifecycle.Name)

	if l == nil {
		lifecycleService.Create(lifecycle)
	} else {
		lifecycle.SetId(l.Id)
		lifecycleService.Update(l.Id, lifecycle)
	}
}

func transformNamesToIds(ids map[string]string, names []string) []string {
	var envIds = make([]string, len(names))
	for index, n := range names {
		envIds[index] = ids[n]
	}
	return envIds
}
