package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/gauravchaudhary999/k8s-octopus/cmd"
	"github.com/gauravchaudhary999/k8s-octopus/model"
)

func main() {
	setFlags()
	switch *cmdPtr {
	case "createOrUpdateSpace":
		smtm := strings.Split(spaceManagersTeamMembers, ",")
		smt := strings.Split(spaceManagersTeams, ",")
		s := model.Space{
			Name:                     spaceName,
			Description:              ifEmptyThen(spaceDescription, spaceName),
			IsDefault:                isDefault,
			TaskQueueStopped:         taskQueueStopped,
			SpaceManagersTeamMembers: smtm,
			SpaceManagersTeams:       smt,
		}
		cmd.CreateOrUpdateSpace(&s)
	case "createOrUpdateEnvironment":
		env := model.Environment{
			Name:                       envName,
			Description:                ifEmptyThen(envDescription, envName),
			SortOrder:                  envSortOrder,
			UseGuidedFailure:           envUseGuidedFailure,
			AllowDynamicInfrastructure: allowDynamicInfrastructure,
		}
		cmd.CreateOrUpdateEnvironment(spaceName, &env)

	case "createOrUpdateAccount":
		acc := model.Account{
			Name:        ifEmptyThen(accName, envName),
			Description: ifEmptyThen(ifEmptyThen(accDescription, accName), envName),
			AccountType: accountType,
			Token: model.AccountToken{
				HasValue: true,
				NewValue: os.Getenv(tokenName),
			},
		}
		cmd.CreateOrUpdateAccount(spaceName, []string{envName}, &acc)

	case "createOrUpdateMachine":
		roles := strings.Split(macRoles, ",")

		machine := model.Machine{
			Name:                            ifEmptyThen(macName, envName),
			Roles:                           roles,
			TenantedDeploymentParticipation: tenantedDeploymentParticipation,
			Endpoint: model.Endpoint{
				CommunicationStyle:  macCommunicationStyle,
				ClusterUrl:          clusterUrl,
				SkipTlsVerification: macSkipTlsVerification,
				Authentication: model.MachineAuthentication{
					AuthenticationType: authenticationType,
				},
			},
		}
		cmd.CreateOrUpdateMachine(spaceName, []string{envName}, ifEmptyThen(accName, envName), &machine)
	case "createOrUpdateLifecycle":
		atm := make(map[string][]string)
		otm := make(map[string][]string)
		at := strings.Split(automaticTargets, ",")
		ot := strings.Split(optionalTargets, ",")

		for _, target := range at {
			atm[target] = []string{target}
		}
		for _, target := range ot {
			otm[target] = []string{target}
		}

		phOrder := strings.Split(phasesOrder, ",")
		lifecycle := model.Lifecycle{
			Name:        lifecycleName,
			Description: ifEmptyThen(lifecycleDescription, lifecycleName),
		}

		cmd.CreateOrUpdateLifecycle(spaceName, phOrder, phOrder, atm, otm, &lifecycle)

	case "createOrUpdateFeed":
		feed := model.Feed{
			Name:     feedName,
			FeedType: feedType,
			FeedUri:  feedUri,
			Username: os.Getenv(feedUserName),
			Password: model.SecretValue{
				HasValue: true,
				NewValue: os.Getenv(feedPassword),
			},
		}
		cmd.CreateOrUpdateFeed(spaceName, &feed)

	case "createOrUpdateActionTemplate":
		cmd.CreateOrUpdateActionTemplate(spaceName, atName, ifEmptyThen(atDescription, atName), actionTemplateType)

	case "configureEnvironment":
		if envName == "" {
			panic("Environment Name is a compulsory field")
		}

		env := model.Environment{
			Name:                       envName,
			Description:                ifEmptyThen(envDescription, envName),
			SortOrder:                  envSortOrder,
			UseGuidedFailure:           envUseGuidedFailure,
			AllowDynamicInfrastructure: allowDynamicInfrastructure,
		}
		acc := model.Account{
			Name:        ifEmptyThen(accName, envName),
			Description: ifEmptyThen(accDescription, envName),
			AccountType: accountType,
			Token: model.AccountToken{
				HasValue: true,
				NewValue: os.Getenv(tokenName),
			},
		}
		roles := strings.Split(macRoles, ",")

		machine := model.Machine{
			Name:                            ifEmptyThen(macName, envName),
			Roles:                           roles,
			TenantedDeploymentParticipation: tenantedDeploymentParticipation,
			Endpoint: model.Endpoint{
				CommunicationStyle:  macCommunicationStyle,
				ClusterUrl:          clusterUrl,
				SkipTlsVerification: macSkipTlsVerification,
				Authentication: model.MachineAuthentication{
					AuthenticationType: authenticationType,
				},
			},
		}
		cmd.ConfigureEnvironment(spaceName, &env, &acc, &machine)

	case "configureOctopus":
		configureOctopus()
	case "createOrUpdateProject":
		p := model.Project{
			Name:        projectName,
			Description: ifEmptyThen(projectDescription, projectName),
			IsDisabled:  isProjectDisabled,
		}
		cmd.CreateOrUpdateProject(spaceName, projectGroupName, lifecycleName, &p)
	case "createOrUpdateProjectGroup":
		pg := model.ProjectGroup{
			Name:        projectGroupName,
			Description: ifEmptyThen(projectGroupDescription, projectGroupName),
		}

		cmd.CreateOrUpdateProjectGroup(spaceName, &pg)

	case "updateDeploymentProcess":
		atNameArray := strings.Split(atNameArr, ",")
		cmd.UpdateDeploymentProcess(spaceName, projectName, atNameArray)
	case "updateVariableSet":
		cmd.UpdateVariableSet(spaceName, projectName, filepath, imageTag)
	case "uploadPackage":
		cmd.UploadPackage(spaceName, filepath)
	case "createRelease":
		cmd.CreateRelease(spaceName, projectName, serviceChartVersion, helmClientVersion, servicePackageVersion, apiGatewayToolVersion)
	default:
		panic("Command not supported")

	}
}

func configureOctopus() {
	smtm := strings.Split(spaceManagersTeamMembers, ",")
	smt := strings.Split(spaceManagersTeams, ",")
	s := model.Space{
		Name:                     spaceName,
		Description:              ifEmptyThen(spaceDescription, spaceName),
		IsDefault:                isDefault,
		TaskQueueStopped:         taskQueueStopped,
		SpaceManagersTeamMembers: smtm,
		SpaceManagersTeams:       smt,
	}

	//Environment Variables
	envNameArray := strings.Split(envNameArr, ",")
	envDescriptionArray := strings.Split(envDescriptionArr, ",")
	envSortOrderArray := strings.Split(envSortOrderArr, ",")
	envUseGuidedFailureArray := strings.Split(envUseGuidedFailureArr, ",")
	allowDynamicInfrastructureArray := strings.Split(allowDynamicInfrastructureArr, ",")

	//Account Variables
	accNameArray := strings.Split(accNameArr, ",")
	accDescriptionArray := strings.Split(accDescriptionArr, ",")
	accountTypeArray := strings.Split(accountTypeArr, ",")
	tokenNameArray := strings.Split(tokenNameArr, ",")

	//Machine variables
	macNameArray := strings.Split(macNameArr, ",")
	macRolesArray := strings.Split(macRolesArr, ",")
	tenantedDeploymentParticipationArray := strings.Split(tenantedDeploymentParticipationArr, ",")
	macCommunicationStyleArray := strings.Split(macCommunicationStyleArr, ",")
	clusterUrlArray := strings.Split(clusterUrlArr, ",")
	authenticationTypeArray := strings.Split(authenticationTypeArr, ",")
	macSkipTlsVerificationArray := strings.Split(macSkipTlsVerificationArr, ",")

	//Action Templates
	atNameArray := strings.Split(atNameArr, ",")
	atDescriptionArray := strings.Split(atDescriptionArr, ",")
	actionTemplateTypeArray := strings.Split(actionTemplateTypeArr, ",")

	var actionTemplateInputs = make([]model.ActionTemplateInput, len(envNameArray))
	var octopusEnvironments = make([]model.OctopusEnvironment, len(envNameArray))

	for i, _ := range envNameArray {
		environmentName := valueFromArray(envNameArray, i, "")
		sortOrder, _ := strconv.Atoi(valueFromArray(envSortOrderArray, i, "0"))
		useGuidedFailure, _ := strconv.ParseBool(valueFromArray(envUseGuidedFailureArray, i, "true"))
		allowDynamicInfrastructure, _ := strconv.ParseBool(valueFromArray(allowDynamicInfrastructureArray, i, "true"))

		env := model.Environment{
			Name:                       environmentName,
			Description:                valueFromArray(envDescriptionArray, i, environmentName),
			SortOrder:                  sortOrder,
			UseGuidedFailure:           useGuidedFailure,
			AllowDynamicInfrastructure: allowDynamicInfrastructure,
		}
		acc := model.Account{
			Name:        valueFromArray(accNameArray, i, environmentName),
			Description: valueFromArray(accDescriptionArray, i, environmentName),
			AccountType: valueFromArray(accountTypeArray, i, "Token"),
			Token: model.AccountToken{
				HasValue: true,
				NewValue: os.Getenv(valueFromArray(tokenNameArray, i, "TOKEN")),
			},
		}

		macineRoles := valueFromArray(macRolesArray, i, "admin")
		roles := strings.Split(macineRoles, ",")
		machineSkipTlsVerification, _ := strconv.ParseBool(valueFromArray(macSkipTlsVerificationArray, i, "true"))

		machine := model.Machine{
			Name:                            valueFromArray(macNameArray, i, environmentName),
			Roles:                           roles,
			TenantedDeploymentParticipation: valueFromArray(tenantedDeploymentParticipationArray, i, "TenantedOrUntenanted"),
			Endpoint: model.Endpoint{
				CommunicationStyle:  valueFromArray(macCommunicationStyleArray, i, "Kubernetes"),
				ClusterUrl:          valueFromArray(clusterUrlArray, i, ""),
				SkipTlsVerification: machineSkipTlsVerification,
				Authentication: model.MachineAuthentication{
					AuthenticationType: valueFromArray(authenticationTypeArray, i, "KubernetesStandard"),
				},
			},
		}
		octopusEnvironments[i] = model.OctopusEnvironment{Environment: &env, Account: &acc, Machine: &machine}
		actionTemplateInputs[i] = model.ActionTemplateInput{Name: valueFromArray(atNameArray, i, ""), Desc: valueFromArray(atDescriptionArray, i, ""), Type: valueFromArray(actionTemplateTypeArray, i, "")}
	}

	atm := make(map[string][]string)
	otm := make(map[string][]string)
	at := strings.Split(automaticTargets, ",")
	ot := strings.Split(optionalTargets, ",")

	for _, target := range at {
		atm[target] = []string{target}
	}
	for _, target := range ot {
		otm[target] = []string{target}
	}
	phOrder := strings.Split(phasesOrder, ",")
	lifecycle := model.Lifecycle{
		Name:        lifecycleName,
		Description: ifEmptyThen(lifecycleDescription, lifecycleName),
	}

	feed := model.Feed{
		Name:     feedName,
		FeedType: feedType,
		FeedUri:  feedUri,
		Username: os.Getenv(feedUserName),
		Password: model.SecretValue{
			HasValue: true,
			NewValue: os.Getenv(feedPassword),
		},
	}

	cmd.ConfigureOctopus(&s, octopusEnvironments, phOrder, phOrder, atm, otm, &lifecycle, &feed, actionTemplateInputs)
}

func ifEmptyThen(name, otherName string) string {
	if name == "" {
		return otherName
	} else {
		return name
	}
}

func valueFromArray(array []string, index int, defaultValue string) string {
	if len(array) == 0 {
		return defaultValue
	} else if len(array) == 1 && array[0] == "" {
		return defaultValue
	} else {
		return array[index]
	}
}
