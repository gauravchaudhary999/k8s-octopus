package octopus

import (
	"fmt"
)

type RawPackageService struct {
	host  string
	token string
}

func NewRawPackageService(host string, token string) *RawPackageService {
	return &RawPackageService{host: host, token: token}
}

func (ps RawPackageService) Create(spaceId, filepath string) {
	url := fmt.Sprintf("/api/%s/packages/raw", spaceId)
	newClient(ps.host, ps.token, url, POST, filepath, nil).executeRaw()
}
