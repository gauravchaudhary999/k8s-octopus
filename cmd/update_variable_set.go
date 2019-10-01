package cmd

import (
	"io/ioutil"
	"log"

	"github.com/gauravchaudhary999/k8s-octopus/model"
	"github.com/gauravchaudhary999/k8s-octopus/octopus"
	"gopkg.in/yaml.v2"
)

type EnvDataBlock struct {
	Env  string              `yaml:"env"`
	Data []map[string]string `yaml:"data"`
}
type VariableFile struct {
	EnvVariables []EnvDataBlock `yaml:"env_variables"`
}

func UpdateVariableSet(spaceName, projectName, filePath, imageTag string) {

	spaceService := octopus.NewSpaceService(octopusHost, octopusApiKey)
	projectService := octopus.NewProjectService(octopusHost, octopusApiKey)
	variableSetService := octopus.NewVariableSetService(octopusHost, octopusApiKey)
	environmentService := octopus.NewEnvironmentService(octopusHost, octopusApiKey)

	space := spaceService.GetByName(spaceName)
	if space == nil {
		panic("Space Not Found")
	}
	spaceId := space.Id
	project := projectService.GetByName(spaceId, projectName)

	if project == nil {
		panic("Project Not Found")
	}

	envs := environmentService.GetAll(spaceId)
	envArray := make([]model.IdNamePair, len(envs.Items))
	for i, e := range envs.Items {
		envArray[i] = model.IdNamePair{
			Id:   e.Id,
			Name: e.Name,
		}
	}
	projectId := project.Id
	variableSet := variableSetService.Get(spaceId, project.VariableSetId)
	variableSetId := variableSet.Id
	variableSetVersion := variableSet.Version

	var variable VariableFile
	variable.read(filePath)

	var variables []model.Variable
	for _, e := range variable.EnvVariables {
		environment := environmentService.GetByName(spaceId, e.Env)

		for _, v := range e.Data {

			for key, value := range v {
				if environment == nil {
					v := model.Variable{
						Name:        key,
						Value:       value,
						Description: key,
						Type:        "string",
						IsSensitive: false,
						IsEditable:  true,
					}
					variables = append(variables, v)
				} else {
					v := model.Variable{
						Name:        key,
						Value:       value,
						Description: key,
						Type:        "string",
						IsSensitive: false,
						IsEditable:  true,
						Scope: model.Scope{
							Environment: []string{environment.Id},
						},
					}
					variables = append(variables, v)
				}
			}
		}
	}
	v := model.Variable{
		Name:        "artifact_version_latest",
		Value:       imageTag,
		Description: "artifact_version_latest",
		Type:        "string",
		IsSensitive: false,
		IsEditable:  true,
	}
	variables = append(variables, v)

	roles := make([]model.IdNamePair, 1)
	roles[0] = model.IdNamePair{
		Id:   "admin",
		Name: "admin",
	}

	vSet := model.VariableSet{
		Id:        variableSetId,
		SpaceId:   spaceId,
		OwnerId:   projectId,
		Version:   variableSetVersion,
		Variables: variables,
		ScopeValues: model.ScopeValues{
			Environments: envArray,
			Roles:        roles,
		},
	}

	variableSetService.Update(variableSetId, &vSet)
}

func (v *VariableFile) read(filePath string) *VariableFile {

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, v)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return v
}
