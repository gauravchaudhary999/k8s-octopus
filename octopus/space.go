package octopus

import (
	"fmt"

	"github.com/gauravchaudhary999/k8s-octopus/model"
)

type SpaceService struct {
	host  string
	token string
}

func NewSpaceService(host string, token string) *SpaceService {
	return &SpaceService{host: host, token: token}
}

func (ss SpaceService) Get(spaceId string) *model.Space {
	space := new(model.Space)
	url := fmt.Sprintf("/api/spaces/%s", spaceId)
	newClient(ss.host, ss.token, url, GET, nil, space).execute()
	return space
}

func (ss SpaceService) Create(s *model.Space) *model.Space {
	space := new(model.Space)
	newClient(ss.host, ss.token, "/api/spaces", POST, *s, space).execute()
	return space
}

func (ss SpaceService) Update(spaceId string, space *model.Space) *model.Space {
	space1 := new(model.Space)
	url := fmt.Sprintf("/api/spaces/%s", spaceId)
	newClient(ss.host, ss.token, url, PUT, *space, space1).execute()
	return space1
}

func (ss SpaceService) GetAll() *model.Spaces {
	spaces := new(model.Spaces)
	newClient(ss.host, ss.token, "/api/spaces", GET, nil, spaces).execute()
	return spaces
}

func (ss SpaceService) GetByName(name string) *model.Space {
	spaces := ss.GetAll()
	return spaces.GetSpace(name)
}

func (ss SpaceService) Delete(spaceId string) {
	url := fmt.Sprintf("/api/spaces/%s", spaceId)
	newClient(ss.host, ss.token, url, DELETE, nil, nil).execute()
}
