package cmd

import "github.com/gauravchaudhary999/k8s-octopus/octopus"

func UploadPackage(spaceName string, filename string) {
	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	rawPackageService := octopus.NewRawPackageService(octopusHost, octopusApiKey)

	space := spaceService.GetByName(spaceName)
	if space == nil {
		panic("Space Not Found")
	}
	spaceId := space.Id

	rawPackageService.Create(spaceId, filename)
}
