package main

import (
	"flag"
)

//Space
var spaceName, spaceDescription, spaceManagersTeamMembers, spaceManagersTeams string
var taskQueueStopped, isDefault bool

//Environment
var envName, envDescription string
var envSortOrder int
var envUseGuidedFailure, allowDynamicInfrastructure bool

//Account
var accName, accDescription, accountType, tokenName string

//Machine
var macName, tenantedDeploymentParticipation, macCommunicationStyle, clusterUrl, authenticationType, macRoles string
var macSkipTlsVerification bool

//Lifecycle
var lifecycleName, lifecycleDescription, phasesOrder, automaticTargets, optionalTargets string

//Feed
var feedName, feedType, feedUri, feedUserName, feedPassword string
var cmdPtr *string

//ActionTemplate
var atName, atDescription, actionTemplateType string

//Configure Octopus
var envNameArr, envDescriptionArr string
var envSortOrderArr string
var envUseGuidedFailureArr, allowDynamicInfrastructureArr string
var accNameArr, accDescriptionArr, accountTypeArr, tokenNameArr string
var macNameArr, tenantedDeploymentParticipationArr, macCommunicationStyleArr, clusterUrlArr, authenticationTypeArr, macRolesArr string
var macSkipTlsVerificationArr string
var atNameArr, atDescriptionArr, actionTemplateTypeArr string

//Project
var projectName, projectDescription string
var isProjectDisabled bool

//ProjectGroup
var projectGroupName, projectGroupDescription string

//VariableSet
var filepath, imageTag string

//Release
var serviceChartVersion, helmClientVersion, servicePackageVersion, apiGatewayToolVersion string

func setFlags() {
	cmdPtr = flag.String("cmd", "", "valid command name from createOrUpdateSpace")
	// Space flags
	flag.StringVar(&spaceName, "spaceName", "", "Space Name")
	flag.StringVar(&spaceDescription, "spaceDescription", "", "Space Description")
	flag.StringVar(&spaceManagersTeamMembers, "spaceManagersTeamMembers", "teams-managers", "Space Managers Team Member")
	flag.StringVar(&spaceManagersTeams, "spaceManagersTeams", "teams-managers", "Space Managers Teams")
	flag.BoolVar(&taskQueueStopped, "taskQueueStopped", true, "TaskQueueStopped")
	flag.BoolVar(&isDefault, "isDefault", false, "IsDefault")

	// Environment flags
	flag.StringVar(&envName, "envName", "", "Environment Name")
	flag.StringVar(&envDescription, "envDescription", "", "Environment Description")
	flag.IntVar(&envSortOrder, "envSortOrder", 0, "Environment Sort Order")
	flag.BoolVar(&envUseGuidedFailure, "envUseGuidedFailure", true, "UseGuidedFailure")
	flag.BoolVar(&allowDynamicInfrastructure, "allowDynamicInfrastructure", true, "AllowDynamicInfrastructure")

	//Account
	flag.StringVar(&accName, "accName", "", "Account Name")
	flag.StringVar(&accDescription, "accDescription", "", "Account Description")
	flag.StringVar(&accountType, "accountType", "Token", "Account Type")
	flag.StringVar(&tokenName, "tokenName", "ACCOUNT_TOKEN", "Token Environment Variable Name")

	//Machine
	flag.StringVar(&macName, "macName", "", "Machine Name")
	flag.StringVar(&macRoles, "macRoles", "admin", "Roles")
	flag.StringVar(&tenantedDeploymentParticipation, "tenantedDeploymentParticipation", "TenantedOrUntenanted", "Machine TenantedDeploymentParticipation")
	flag.StringVar(&macCommunicationStyle, "macCommunicationStyle", "Kubernetes", "Machine Communication Style")
	flag.StringVar(&clusterUrl, "clusterUrl", "", "Cluster Url")
	flag.StringVar(&authenticationType, "authenticationType", "KubernetesStandard", "Authentication Type")
	flag.BoolVar(&macSkipTlsVerification, "macSkipTlsVerification", true, "Machine SkipTlsVerification")

	//Lifecycle
	flag.StringVar(&lifecycleName, "lifecycleName", "", "Lifecycle Name")
	flag.StringVar(&lifecycleDescription, "lifecycleDescription", "", "Lifecycle Description")
	flag.StringVar(&phasesOrder, "phasesOrder", "", "Lifecycle Phases in order of executon")
	flag.StringVar(&automaticTargets, "automaticTargets", "", "Lifecycle AutomaticDeploymentTargets")
	flag.StringVar(&optionalTargets, "optionalTargets", "", "Lifecycle OptionalDeploymentTargets")

	//Feed
	flag.StringVar(&feedName, "feedName", "", "Feed Name")
	flag.StringVar(&feedType, "feedType", "Helm", "Feed Type")
	flag.StringVar(&feedUri, "feedUri", "", "Feed Uri")
	flag.StringVar(&feedUserName, "feedUserName", "FEED_USERNAME", "Feed Username env variable")
	flag.StringVar(&feedPassword, "feedPassword", "FEED_PASSWORD", "Feed Password env variable")

	//Action Template
	flag.StringVar(&atName, "atName", "", "Action Template Name")
	flag.StringVar(&atDescription, "atDescription", "", "Action Template Description")
	flag.StringVar(&actionTemplateType, "actionTemplateType", "", "Action Template Type")

	//Configure Octopus
	flag.StringVar(&envNameArr, "envNameArr", "", "Environment Name")
	flag.StringVar(&envDescriptionArr, "envDescriptionArr", "", "Environment Description")
	flag.StringVar(&envSortOrderArr, "envSortOrderArr", "", "Environment Sort Order")
	flag.StringVar(&envUseGuidedFailureArr, "envUseGuidedFailureArr", "", "UseGuidedFailure")
	flag.StringVar(&allowDynamicInfrastructureArr, "allowDynamicInfrastructureArr", "", "AllowDynamicInfrastructure")

	flag.StringVar(&accNameArr, "accNameArr", "", "Account Name")
	flag.StringVar(&accDescriptionArr, "accDescriptionArr", "", "Account Description")
	flag.StringVar(&accountTypeArr, "accountTypeArr", "", "Account Type")
	flag.StringVar(&tokenNameArr, "tokenNameArr", "", "Token Environment Variable Name")

	flag.StringVar(&macNameArr, "macNameArr", "", "Machine Name")
	flag.StringVar(&macRolesArr, "macRolesArr", "", "Roles")
	flag.StringVar(&tenantedDeploymentParticipationArr, "tenantedDeploymentParticipationArr", "", "Machine TenantedDeploymentParticipation")
	flag.StringVar(&macCommunicationStyleArr, "macCommunicationStyleArr", "", "Machine Communication Style")
	flag.StringVar(&clusterUrlArr, "clusterUrlArr", "", "Cluster Url")
	flag.StringVar(&authenticationTypeArr, "authenticationTypeArr", "", "Authentication Type")
	flag.StringVar(&macSkipTlsVerificationArr, "macSkipTlsVerificationArr", "", "Machine SkipTlsVerification")

	flag.StringVar(&atNameArr, "atNameArr", "", "Action Template Name")
	flag.StringVar(&atDescriptionArr, "atDescriptionArr", "", "Action Template Description")
	flag.StringVar(&actionTemplateTypeArr, "actionTemplateTypeArr", "", "Action Template Type")

	// Project
	flag.StringVar(&projectName, "projectName", "", "Project Name")
	flag.StringVar(&projectDescription, "projectDescription", "", "Project Description")
	flag.BoolVar(&isProjectDisabled, "isProjectDisabled", false, "IsProjectDisabled")

	// Project Variables
	flag.StringVar(&projectGroupName, "projectGroupName", "", "Project Group Name")
	flag.StringVar(&projectGroupDescription, "projectGroupDescription", "", "Project Group Description")

	//VariableSet
	flag.StringVar(&filepath, "filepath", "", "Filepath of the yaml file containing variables")
	flag.StringVar(&imageTag, "imageTag", "latest", "Docker Image Tag")

	//Release
	flag.StringVar(&serviceChartVersion, "serviceChartVersion", "", "Helm Chart Version for the service")
	flag.StringVar(&helmClientVersion, "helmClientVersion", "", "Helm client Version")
	flag.StringVar(&servicePackageVersion, "servicePackageVersion", "", "Package Version for the service")
	flag.StringVar(&apiGatewayToolVersion, "apiGatewayToolVersion", "", "Api gateway tool version")

	flag.Parse()
}
