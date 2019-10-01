package model

type Project struct {
	Id, ClonedFromProjectId, SpaceId, Name, Description             string
	ProjectGroupId, LifecycleId, DeploymentProcessId, VariableSetId string
	IsDisabled                                                      bool
}

type Projects struct {
	Items []Project
}

func (s Projects) GetProject(name string) *Project {
	for _, project := range s.Items {
		if project.Name == name {
			return &project
		}
	}
	return nil
}

func (p *Project) SetId(id string) {
	p.Id = id
}

func (p *Project) SetSpaceId(spaceId string) {
	p.SpaceId = spaceId
}

func (p *Project) SetProjectGroupId(projectGroupId string) {
	p.ProjectGroupId = projectGroupId
}

func (p *Project) SetLifecycleId(lifecycleId string) {
	p.LifecycleId = lifecycleId
}
