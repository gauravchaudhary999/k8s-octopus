package model

type ActionTemplatePackagesProperties struct {
	ValuesFilePath         string `json:"ValuesFilePath,omitempty"`
	PerformVariableReplace string `json:"PerformVariableReplace,omitempty"`
	Extract                string `json:"Extract,omitempty"`
}

type ActionTemplatePackages struct {
	Id, Name, PackageId, FeedId, AcquisitionLocation string
	Properties                                       ActionTemplatePackagesProperties
}

type ActionTemplate struct {
	Id, Name, Description, ActionType, SpaceId string
	Packages                                   []ActionTemplatePackages
	Properties                                 map[string]string
	Version                                    int
}

type ActionTemplates struct {
	Items []ActionTemplate
}

func (s ActionTemplates) GetActionTemplate(name string) *ActionTemplate {
	for _, actiontemplate := range s.Items {
		if actiontemplate.Name == name {
			return &actiontemplate
		}
	}
	return nil
}

func (at *ActionTemplate) SetId(id string) {
	at.Id = id
}

func (at *ActionTemplate) SetSpaceId(spaceId string) {
	at.SpaceId = spaceId
}

func (atp *ActionTemplatePackages) SetFeedId(feedId string) {
	atp.FeedId = feedId
}

func (at *ActionTemplate) AddProperty(key, value string) {
	at.Properties[key] = value
}

func (at *ActionTemplate) SetProperties(properties map[string]string) {
	at.Properties = properties
}
