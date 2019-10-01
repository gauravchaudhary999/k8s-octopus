package model

type SecretValue struct {
	HasValue bool   `json:"HasValue"`
	NewValue string `json:"NewValue"`
}

type Feed struct {
	Id       string      `json:"Id"`
	Name     string      `json:"Name"`
	SpaceId  string      `json:"SpaceId"`
	FeedType string      `json:"FeedType"`
	FeedUri  string      `json:"FeedUri"`
	Username string      `json:"Username"`
	Password SecretValue `json:"Password"`
}

type Feeds struct {
	Items []Feed
}

func (feeds Feeds) GetFeed(name string) *Feed {
	for _, feed := range feeds.Items {
		if feed.Name == name {
			return &feed
		}
	}
	return nil
}

func (feeds Feeds) GetFeedByType(feedType string) *Feed {
	for _, feed := range feeds.Items {
		if feed.FeedType == feedType {
			return &feed
		}
	}
	return nil
}

func (feed *Feed) SetSpaceId(spaceId string) {
	feed.SpaceId = spaceId
}

func (feed *Feed) SetId(id string) {
	feed.Id = id
}
