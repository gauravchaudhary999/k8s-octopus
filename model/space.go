package model

type Spaces struct {
	Items []Space
}

func (s Spaces) GetSpace(name string) *Space {
	for _, space := range s.Items {
		if space.Name == name {
			return &space
		}
	}
	return nil
}

type Space struct {
	Id, Name, Description       string
	IsDefault, TaskQueueStopped bool
	SpaceManagersTeamMembers    []string
	SpaceManagersTeams          []string
}

func (space *Space) SetId(id string) {
	space.Id = id
}

func (space *Space) UpdateSpaceManaget(managers []string) {
	for _, manager := range managers {
		if !space.exists(manager) {
			space.SpaceManagersTeams = append(space.SpaceManagersTeams, manager)
		}
	}
}

func (space Space) exists(manager string) bool {

	for _, existing := range space.SpaceManagersTeams {
		if existing == manager {
			return true
		}
	}
	return false
}
