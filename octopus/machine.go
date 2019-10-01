package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type MachineService struct {
	host  string
	token string
}

func NewMachineService(host string, token string) *MachineService {
	return &MachineService{host: host, token: token}
}
func (ms MachineService) Create(e *model.Machine) *model.Machine {
	machine := new(model.Machine)
	url := fmt.Sprintf("/api/%s/machines", e.SpaceId)
	newClient(ms.host, ms.token, url, POST, *e, machine).execute()
	return machine
}

func (ms MachineService) Update(machineId string, e *model.Machine) *model.Machine {
	machine := new(model.Machine)
	url := fmt.Sprintf("/api/%s/machines/%s", e.SpaceId, machineId)
	newClient(ms.host, ms.token, url, PUT, *e, machine).execute()
	return machine
}

func (ms MachineService) Get(spaceId string, machineId string) *model.Machine {
	machine := new(model.Machine)
	url := fmt.Sprintf("/api/%s/machines/%s", spaceId, machineId)
	newClient(ms.host, ms.token, url, GET, nil, machine).execute()
	return machine
}

func (ms MachineService) GetAll(spaceId string) *model.Machines {
	machines := new(model.Machines)
	url := fmt.Sprintf("/api/%s/machines", spaceId)
	newClient(ms.host, ms.token, url, GET, nil, machines).execute()
	return machines
}

func (ms MachineService) GetByName(spaceId string, name string) *model.Machine {
	machines := ms.GetAll(spaceId)
	return machines.GetMachine(name)
}
