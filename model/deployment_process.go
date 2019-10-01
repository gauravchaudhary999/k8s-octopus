package model

type Action struct {
	Id, Name, ActionType, WorkerPoolId                       string
	Environments, ExcludedEnvironments, Channels, TenantTags []string
	IsDisabled, CanBeUsedForProjectVersioning, IsRequired    bool
	Properties                                               map[string]string
	Packages                                                 []ActionTemplatePackages
}

type Step struct {
	Id, Name, PackageRequirement, Condition, StartTrigger string
	Properties                                            map[string]string
	Actions                                               []Action
}

type DeploymentProcess struct {
	Id, ProjectId, LastSnapshotId, SpaceId string
	Version                                int
	Steps                                  []Step
}

type DeploymentProcesses struct {
	Items []DeploymentProcess
}
