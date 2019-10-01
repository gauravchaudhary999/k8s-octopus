package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateActionTemplate(spaceName, name, description, actionTemplateType string) {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	feedService := octopus.NewFeedService(octopusHost, octopusApiKey)
	actionTemplateService := octopus.NewActionTemplateService(octopusHost, octopusApiKey)

	s := spaceService.GetByName(spaceName)
	if s == nil {
		panic("Space does not exist")
	}
	createOrUpdateActionTemplate(s.Id, name, description, actionTemplateType, feedService, actionTemplateService)
}

func createOrUpdateActionTemplate(spaceId, name, description, actionTemplateType string, feedService *octopus.FeedService, actionTemplateService *octopus.ActionTemplateService) {
	f := feedService.GetBuiltIn(spaceId)
	if f == nil {
		panic("Built In feed not found")
	}

	var template *model.ActionTemplate

	switch actionTemplateType {
	case "Helm":
		template = getHelmUpgradeActionTemplate(name, description, spaceId, f.Id)
	case "Kong":
		template = getKongActionTemplate(name, description, spaceId, f.Id)
	default:
		panic("Template not supported")
	}

	at := actionTemplateService.GetByName(spaceId, name)

	if at == nil {
		actionTemplateService.Create(template)
	} else {
		template.SetId(at.Id)
		actionTemplateService.Update(at.Id, template)
	}
}
