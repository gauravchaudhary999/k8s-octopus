package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type VariableSetService struct {
	host  string
	token string
}

func NewVariableSetService(host string, token string) *VariableSetService {
	return &VariableSetService{host: host, token: token}
}

func (ps VariableSetService) Update(variableId string, e *model.VariableSet) *model.VariableSet {
	variable := new(model.VariableSet)
	url := fmt.Sprintf("/api/%s/variables/%s", e.SpaceId, variableId)
	newClient(ps.host, ps.token, url, PUT, *e, variable).execute()
	return variable
}

func (ps VariableSetService) Get(spaceId string, variableId string) *model.VariableSet {
	variable := new(model.VariableSet)
	url := fmt.Sprintf("/api/%s/variables/%s", spaceId, variableId)
	newClient(ps.host, ps.token, url, GET, nil, variable).execute()
	return variable
}

func (ps VariableSetService) GetAll(spaceId string) *model.VariableSet {
	variable := new(model.VariableSet)
	url := fmt.Sprintf("/api/%s/variables/all", spaceId)
	newClient(ps.host, ps.token, url, GET, nil, variable).execute()
	return variable
}
