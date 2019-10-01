package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type ActionTemplateService struct {
	host  string
	token string
}

func NewActionTemplateService(host string, token string) *ActionTemplateService {
	return &ActionTemplateService{host: host, token: token}
}

func (ats ActionTemplateService) Create(e *model.ActionTemplate) *model.ActionTemplate {
	actiontemplate := new(model.ActionTemplate)
	url := fmt.Sprintf("/api/%s/actiontemplates", e.SpaceId)
	newClient(ats.host, ats.token, url, POST, *e, actiontemplate).execute()
	return actiontemplate
}

func (ats ActionTemplateService) Update(actionTemplateId string, e *model.ActionTemplate) *model.ActionTemplate {
	actiontemplate := new(model.ActionTemplate)
	url := fmt.Sprintf("/api/%s/actiontemplates/%s", e.SpaceId, actionTemplateId)
	newClient(ats.host, ats.token, url, PUT, *e, actiontemplate).execute()
	return actiontemplate
}

func (ats ActionTemplateService) Get(spaceId string, actionTemplateId string) *model.ActionTemplate {
	actiontemplate := new(model.ActionTemplate)
	url := fmt.Sprintf("/api/%s/actiontemplates/%s", spaceId, actionTemplateId)
	newClient(ats.host, ats.token, url, GET, nil, actiontemplate).execute()
	return actiontemplate
}

func (ats ActionTemplateService) GetAll(spaceId string) *model.ActionTemplates {
	actiontemplates := new(model.ActionTemplates)
	url := fmt.Sprintf("/api/%s/actiontemplates", spaceId)
	newClient(ats.host, ats.token, url, GET, nil, actiontemplates).execute()
	return actiontemplates
}

func (ats ActionTemplateService) GetByName(spaceId string, name string) *model.ActionTemplate {
	actiontemplates := ats.GetAll(spaceId)
	return actiontemplates.GetActionTemplate(name)
}
