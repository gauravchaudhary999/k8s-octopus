package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type AccountService struct {
	host  string
	token string
}

func NewAccountService(host string, token string) *AccountService {
	return &AccountService{host: host, token: token}
}

func (ac AccountService) Create(a *model.Account) *model.Account {
	acc := new(model.Account)
	url := fmt.Sprintf("/api/%s/accounts", a.SpaceId)
	newClient(ac.host, ac.token, url, POST, *a, acc).execute()
	return acc
}

func (ac AccountService) Update(accountId string, a *model.Account) *model.Account {
	acc := new(model.Account)
	url := fmt.Sprintf("/api/%s/accounts/%s", a.SpaceId, accountId)
	newClient(ac.host, ac.token, url, PUT, *a, acc).execute()
	return acc
}

func (ac AccountService) Get(spaceId string, accountId string) *model.Account {
	acc := new(model.Account)
	url := fmt.Sprintf("/api/%s/account/%s", spaceId, accountId)
	newClient(ac.host, ac.token, url, GET, nil, acc).execute()
	return acc
}

func (ac AccountService) GetAll(spaceId string) *model.Accounts {
	acc := new(model.Accounts)
	url := fmt.Sprintf("/api/%s/accounts", spaceId)
	newClient(ac.host, ac.token, url, GET, nil, acc).execute()
	return acc
}

func (ac AccountService) GetByName(spaceId string, name string) *model.Account {
	accounts := ac.GetAll(spaceId)
	return accounts.GetAccount(name)
}
