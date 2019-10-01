package model

type SelectedPackage struct {
	StepName, ActionName, Version, PackageReferenceName string
}

type Release struct {
	Id, Assembled, ReleaseNotes, ProjectId               string
	ChannelId, ProjectVariableSetSnapshotId              string
	ProjectDeploymentProcessSnapshotId, SpaceId, Version string
	LibraryVariableSetSnapshotIds                        []string
	IgnoreChannelRules                                   bool
	SelectedPackages                                     []SelectedPackage
}

type Releases struct {
	Items []Release
}
