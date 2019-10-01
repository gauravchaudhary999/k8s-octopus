package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type FeedService struct {
	host  string
	token string
}

func NewFeedService(host string, token string) *FeedService {
	return &FeedService{host: host, token: token}
}

func (fs FeedService) Create(e *model.Feed) *model.Feed {
	feed := new(model.Feed)
	url := fmt.Sprintf("/api/%s/feeds", e.SpaceId)
	newClient(fs.host, fs.token, url, POST, *e, feed).execute()
	return feed
}

func (fs FeedService) Update(feedId string, e *model.Feed) *model.Feed {
	feed := new(model.Feed)
	url := fmt.Sprintf("/api/%s/feeds/%s", e.SpaceId, feedId)
	newClient(fs.host, fs.token, url, PUT, *e, feed).execute()
	return feed
}

func (fs FeedService) Get(spaceId string, feedId string) *model.Feed {
	feed := new(model.Feed)
	url := fmt.Sprintf("/api/%s/feeds/%s", spaceId, feedId)
	newClient(fs.host, fs.token, url, GET, nil, feed).execute()
	return feed
}

func (fs FeedService) GetAll(spaceId string) *model.Feeds {
	feed := new(model.Feeds)
	url := fmt.Sprintf("/api/%s/feeds", spaceId)
	newClient(fs.host, fs.token, url, GET, nil, feed).execute()
	return feed
}

func (fs FeedService) GetByName(spaceId string, name string) *model.Feed {
	feeds := fs.GetAll(spaceId)
	return feeds.GetFeed(name)
}

func (fs FeedService) GetBuiltIn(spaceId string) *model.Feed {
	feeds := fs.GetAll(spaceId)
	return feeds.GetFeedByType("BuiltIn")
}
