package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type DeploymentProcessService struct {
	host  string
	token string
}

func NewDeploymentProcessService(host string, token string) *DeploymentProcessService {
	return &DeploymentProcessService{host: host, token: token}
}

func (ps DeploymentProcessService) Update(deploymentProcessId string, e *model.DeploymentProcess) *model.DeploymentProcess {
	deploymentProcess := new(model.DeploymentProcess)
	url := fmt.Sprintf("/api/%s/deploymentprocesses/%s", e.SpaceId, deploymentProcessId)
	newClient(ps.host, ps.token, url, PUT, *e, deploymentProcess).execute()
	return deploymentProcess
}

func (ps DeploymentProcessService) Get(spaceId string, deploymentProcessId string) *model.DeploymentProcess {
	deploymentProcess := new(model.DeploymentProcess)
	url := fmt.Sprintf("/api/%s/deploymentprocesses/%s", spaceId, deploymentProcessId)
	newClient(ps.host, ps.token, url, GET, nil, deploymentProcess).execute()
	return deploymentProcess
}

func (ps DeploymentProcessService) GetAll(spaceId string) *model.DeploymentProcesses {
	deploymentProcesses := new(model.DeploymentProcesses)
	url := fmt.Sprintf("/api/%s/deploymentprocesses", spaceId)
	newClient(ps.host, ps.token, url, GET, nil, deploymentProcesses).execute()
	return deploymentProcesses
}
