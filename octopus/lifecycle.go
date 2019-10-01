package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type LifecycleService struct {
	host  string
	token string
}

func NewLifecycleService(host string, token string) *LifecycleService {
	return &LifecycleService{host: host, token: token}
}

func (ls LifecycleService) Create(e *model.Lifecycle) *model.Lifecycle {
	lifecycle := new(model.Lifecycle)
	url := fmt.Sprintf("/api/%s/lifecycles", e.SpaceId)
	newClient(ls.host, ls.token, url, POST, *e, lifecycle).execute()
	return lifecycle
}

func (ls LifecycleService) Update(lifecycleId string, e *model.Lifecycle) *model.Lifecycle {
	lifecycle := new(model.Lifecycle)
	url := fmt.Sprintf("/api/%s/lifecycles/%s", e.SpaceId, lifecycleId)
	newClient(ls.host, ls.token, url, PUT, *e, lifecycle).execute()
	return lifecycle
}

func (ls LifecycleService) Get(spaceId string, lifecycleId string) *model.Lifecycle {
	lifecycle := new(model.Lifecycle)
	url := fmt.Sprintf("/api/%s/lifecycles/%s", spaceId, lifecycleId)
	newClient(ls.host, ls.token, url, GET, nil, lifecycle).execute()
	return lifecycle
}

func (ls LifecycleService) GetAll(spaceId string) *model.Lifecycles {
	lifecycles := new(model.Lifecycles)
	url := fmt.Sprintf("/api/%s/lifecycles", spaceId)
	newClient(ls.host, ls.token, url, GET, nil, lifecycles).execute()
	return lifecycles
}

func (ls LifecycleService) GetByName(spaceId string, name string) *model.Lifecycle {
	lifecycles := ls.GetAll(spaceId)
	return lifecycles.GetLifecycle(name)
}
