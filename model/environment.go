package model

type Environments struct {
	Items []Environment
}

type Environment struct {
	Id, Name, Description, SpaceId               string
	SortOrder                                    int
	UseGuidedFailure, AllowDynamicInfrastructure bool
}

func (envs Environments) GetEnvironment(name string) *Environment {
	for _, env := range envs.Items {
		if env.Name == name {
			return &env
		}
	}
	return nil
}

func (en *Environment) SetId(id string) {
	en.Id = id
}

func (en *Environment) SetSpaceId(spaceId string) {
	en.SpaceId = spaceId
}
