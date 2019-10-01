package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type ReleaseService struct {
	host  string
	token string
}

func NewReleaseService(host string, token string) *ReleaseService {
	return &ReleaseService{host: host, token: token}
}

func (ps ReleaseService) Create(e *model.Release) *model.Release {
	release := new(model.Release)
	url := fmt.Sprintf("/api/%s/releases", e.SpaceId)
	newClient(ps.host, ps.token, url, POST, *e, release).execute()
	return release
}

func (ps ReleaseService) Update(releaseId string, e *model.Release) *model.Release {
	release := new(model.Release)
	url := fmt.Sprintf("/api/%s/releases/%s", e.SpaceId, releaseId)
	newClient(ps.host, ps.token, url, PUT, *e, release).execute()
	return release
}

func (ps ReleaseService) Get(spaceId string, releaseId string) *model.Release {
	release := new(model.Release)
	url := fmt.Sprintf("/api/%s/releases/%s", spaceId, releaseId)
	newClient(ps.host, ps.token, url, GET, nil, release).execute()
	return release
}

func (ps ReleaseService) GetAll(spaceId string) *model.Releases {
	release := new(model.Releases)
	url := fmt.Sprintf("/api/%s/releases", spaceId)
	newClient(ps.host, ps.token, url, GET, nil, release).execute()
	return release
}
