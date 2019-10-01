package cmd

import (
	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
)

func CreateOrUpdateFeed(spaceName string, feed *model.Feed) {

	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	feedService := octopus.NewFeedService(octopusHost, octopusApiKey)
	space := spaceService.GetByName(spaceName)
	if space == nil {
		panic("Space Not Found")
	}

	createOrUpdateFeed(space.Id, feed, feedService)
}

func createOrUpdateFeed(spaceId string, feed *model.Feed, feedService *octopus.FeedService) string {
	feed.SetSpaceId(spaceId)
	f := feedService.GetByName(spaceId, feed.Name)

	if f == nil {
		f = feedService.Create(feed)
	} else {
		feed.SetId(f.Id)
		f = feedService.Update(f.Id, feed)
	}

	return f.Id
}
