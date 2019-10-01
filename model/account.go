package model

type AccountToken struct {
	HasValue bool
	NewValue string
}

type Account struct {
	Id, Name, Description, AccountType, SpaceId string
	EnvironmentIds                              []string
	Token                                       AccountToken
}

type Accounts struct {
	Items []Account
}

func (accounts Accounts) GetAccount(name string) *Account {
	for _, acc := range accounts.Items {
		if acc.Name == name {
			return &acc
		}
	}
	return nil
}

func (en *Account) SetId(id string) {
	en.Id = id
}

func (en *Account) SetSpaceId(spaceId string) {
	en.SpaceId = spaceId
}

func (en *Account) SetEnvironmentIds(envIds []string) {
	en.EnvironmentIds = envIds
}
