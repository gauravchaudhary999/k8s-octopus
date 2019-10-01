package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type ProjectGroupService struct {
	host  string
	token string
}

func NewProjectGroupService(host string, token string) *ProjectGroupService {
	return &ProjectGroupService{host: host, token: token}
}

func (ps ProjectGroupService) Create(e *model.ProjectGroup) *model.ProjectGroup {
	projectgroup := new(model.ProjectGroup)
	url := fmt.Sprintf("/api/%s/projectgroups", e.SpaceId)
	newClient(ps.host, ps.token, url, POST, *e, projectgroup).execute()
	return projectgroup
}

func (ps ProjectGroupService) Update(projectId string, e *model.ProjectGroup) *model.ProjectGroup {
	projectgroup := new(model.ProjectGroup)
	url := fmt.Sprintf("/api/%s/projectgroups/%s", e.SpaceId, projectId)
	newClient(ps.host, ps.token, url, PUT, *e, projectgroup).execute()
	return projectgroup
}

func (ps ProjectGroupService) Get(spaceId string, projectId string) *model.ProjectGroup {
	projectgroup := new(model.ProjectGroup)
	url := fmt.Sprintf("/api/%s/projectgroups/%s", spaceId, projectId)
	newClient(ps.host, ps.token, url, GET, nil, projectgroup).execute()
	return projectgroup
}

func (ps ProjectGroupService) GetAll(spaceId string) *model.ProjectGroups {
	projectgroups := new(model.ProjectGroups)
	url := fmt.Sprintf("/api/%s/projectgroups", spaceId)
	newClient(ps.host, ps.token, url, GET, nil, projectgroups).execute()
	return projectgroups
}

func (ps ProjectGroupService) GetByName(spaceId string, name string) *model.ProjectGroup {
	projectgroups := ps.GetAll(spaceId)
	return projectgroups.GetProjectGroup(name)
}
