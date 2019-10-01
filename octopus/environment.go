package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type EnvironmentService struct {
	host  string
	token string
}

func NewEnvironmentService(host string, token string) *EnvironmentService {
	return &EnvironmentService{host: host, token: token}
}

func (es EnvironmentService) Create(e *model.Environment) *model.Environment {
	en := new(model.Environment)
	url := fmt.Sprintf("/api/%s/environments", e.SpaceId)
	newClient(es.host, es.token, url, POST, *e, en).execute()
	return en
}

func (es EnvironmentService) Update(envId string, e *model.Environment) *model.Environment {
	en := new(model.Environment)
	url := fmt.Sprintf("/api/%s/environments/%s", e.SpaceId, envId)
	newClient(es.host, es.token, url, PUT, *e, en).execute()
	return en
}

func (es EnvironmentService) Get(spaceId string, envId string) *model.Environment {
	en := new(model.Environment)
	url := fmt.Sprintf("/api/%s/environments/%s", spaceId, envId)
	newClient(es.host, es.token, url, GET, nil, en).execute()
	return en
}

func (es EnvironmentService) GetAll(spaceId string) *model.Environments {
	en := new(model.Environments)
	url := fmt.Sprintf("/api/%s/environments", spaceId)
	newClient(es.host, es.token, url, GET, nil, en).execute()
	return en
}

func (es EnvironmentService) GetByName(spaceId string, name string) *model.Environment {
	envs := es.GetAll(spaceId)
	return envs.GetEnvironment(name)
}
