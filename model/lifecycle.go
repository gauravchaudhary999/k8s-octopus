package model

type Phase struct {
	Id                                                    string `json:"Id,omitempty"`
	Name                                                  string
	AutomaticDeploymentTargets, OptionalDeploymentTargets []string
}
type Lifecycle struct {
	Id, Name, Description, SpaceId string
	Phases                         *[]Phase
}

type Lifecycles struct {
	Items []Lifecycle
}

func (s Lifecycles) GetLifecycle(name string) *Lifecycle {
	for _, lifecycle := range s.Items {
		if lifecycle.Name == name {
			return &lifecycle
		}
	}
	return nil
}

func (l *Lifecycle) SetId(id string) {
	l.Id = id
}

func (l *Lifecycle) SetSpaceId(spaceId string) {
	l.SpaceId = spaceId
}

func (l *Lifecycle) SetPhases(phases *[]Phase) {
	l.Phases = phases
}
