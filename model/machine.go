package model

type MachineAuthentication struct {
	AccountId, AuthenticationType, ClientCertificate string
}

type Endpoint struct {
	Id, CommunicationStyle, ClusterCertificate, ClusterUrl   string
	Namespace, ProxyId, DefaultWorkerPoolId, URI, Thumbprint string
	SkipTlsVerification                                      bool
	Authentication                                           MachineAuthentication
}
type Machine struct {
	Id                              string `json:"Id"`
	SpaceId                         string `json:"SpaceId"`
	Name                            string `json:"Name"`
	Status                          string `json:"Status,omitempty"`
	HealthStatus                    string `json:"HealthStatus,omitempty"`
	TenantedDeploymentParticipation string `json:"TenantedDeploymentParticipation,omitempty"`
	TenantIds                       string `json:"TenantIds,omitempty"`
	TenantTags                      string `json:"TenantTags,omitempty"`
	EnvironmentIds, Roles           []string
	Endpoint                        Endpoint
}

type Machines struct {
	Items []Machine
}

func (s Machines) GetMachine(name string) *Machine {
	for _, machine := range s.Items {
		if machine.Name == name {
			return &machine
		}
	}
	return nil
}

func (m *Machine) SetId(id string) {
	m.Id = id
}

func (m *Machine) SetSpaceId(spaceId string) {
	m.SpaceId = spaceId
}

func (m *Machine) SetEnvironmentIds(envIds []string) {
	m.EnvironmentIds = envIds
}

func (m *MachineAuthentication) SetAccountId(accountId string) {
	m.AccountId = accountId
}
