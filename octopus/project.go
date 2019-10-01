package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type ProjectService struct {
	host  string
	token string
}

func NewProjectService(host string, token string) *ProjectService {
	return &ProjectService{host: host, token: token}
}

func (ps ProjectService) Create(e *model.Project) *model.Project {
	project := new(model.Project)
	url := fmt.Sprintf("/api/%s/projects", e.SpaceId)
	newClient(ps.host, ps.token, url, POST, *e, project).execute()
	return project
}

func (ps ProjectService) Update(projectId string, e *model.Project) *model.Project {
	project := new(model.Project)
	url := fmt.Sprintf("/api/%s/projects/%s", e.SpaceId, projectId)
	newClient(ps.host, ps.token, url, PUT, *e, project).execute()
	return project
}

func (ps ProjectService) Get(spaceId string, projectId string) *model.Project {
	project := new(model.Project)
	url := fmt.Sprintf("/api/%s/projects/%s", spaceId, projectId)
	newClient(ps.host, ps.token, url, GET, nil, project).execute()
	return project
}

func (ps ProjectService) GetAll(spaceId string) *model.Projects {
	project := new(model.Projects)
	url := fmt.Sprintf("/api/%s/projects", spaceId)
	newClient(ps.host, ps.token, url, GET, nil, project).execute()
	return project
}

func (ps ProjectService) GetByName(spaceId string, name string) *model.Project {
	projects := ps.GetAll(spaceId)
	return projects.GetProject(name)
}
