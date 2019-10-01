package model

type ProjectGroup struct {
	Id, Name, Description, SpaceId string
	EnvironmentIds                 []string
}

type ProjectGroups struct {
	Items []ProjectGroup
}

func (s ProjectGroups) GetProjectGroup(name string) *ProjectGroup {
	for _, projectgroup := range s.Items {
		if projectgroup.Name == name {
			return &projectgroup
		}
	}
	return nil
}

func (pg *ProjectGroup) SetId(id string) {
	pg.Id = id
}

func (pg *ProjectGroup) SetSpaceId(spaceId string) {
	pg.SpaceId = spaceId
}
