package model

type IdNamePair struct {
	Id, Name string
}

type Scope struct {
	Channel, Private, Machine, TargetRole, Tenant, User, ParentDeployment, TenantTag, Environment, Role, Action, Trigger, Project []string
}
type Variable struct {
	Id, Type, Description, Name, Value, Prompt string
	IsSensitive, IsEditable                    bool
	Scope                                      Scope
}

type ScopeValues struct {
	Environments, Machines, Actions, Roles, Channels, TenantTags []IdNamePair
}
type VariableSet struct {
	Id, OwnerId, SpaceId string
	Version              int
	Variables            []Variable
	ScopeValues          ScopeValues
}
