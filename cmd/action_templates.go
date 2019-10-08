package cmd

import "github.com/gauravchaudhary999/k8s-octopus/model"

func getHelmUpgradeActionTemplate(name, description, spaceId, feedId string) *model.ActionTemplate {
	properties := make(map[string]string)
	properties["Octopus.Action.Helm.CustomHelmExecutable"] = "windows-amd64\\helm.exe"
	properties["Octopus.Action.Helm.ResetValues"] = "True"
	properties["Octopus.Action.Package.DownloadOnTentacle"] = "False"
	properties["Octopus.Action.Package.FeedId"] = "#{helm_feed_id}"
	properties["Octopus.Action.Package.PackageId"] = "#{helm_chart_package_id}"
	properties["Octopus.Action.Helm.ReleaseName"] = "#{helm_release_name}"
	properties["Octopus.Action.Helm.Namespace"] = "#{namespace}"

	at := model.ActionTemplate{
		Name:        name,
		Description: description,
		SpaceId:     spaceId,
		ActionType:  "Octopus.HelmChartUpgrade",
		Packages: []model.ActionTemplatePackages{
			model.ActionTemplatePackages{Name: "",
				PackageId:           "#{helm_chart_package_id}",
				FeedId:              "#{helm_feed_id}",
				AcquisitionLocation: "Server",
			},
			model.ActionTemplatePackages{Name: "ValuesPack-1",
				PackageId:           "#{helm_custom_package}",
				FeedId:              feedId,
				AcquisitionLocation: "ExecutionTarget",
				Properties: model.ActionTemplatePackagesProperties{
					ValuesFilePath: "values.#{Octopus.Environment.Name}.yaml",
				},
			},
			model.ActionTemplatePackages{Name: "HelmExe",
				PackageId:           "helm-v2",
				FeedId:              feedId,
				AcquisitionLocation: "ExecutionTarget",
				Properties: model.ActionTemplatePackagesProperties{
					PerformVariableReplace: "False",
					Extract:                "True",
				},
			},
		},
		Properties: properties,
	}
	return &at
}

func getKongActionTemplate(name, description, spaceId, feedId string) *model.ActionTemplate {
	properties := make(map[string]string)
	properties["Octopus.Action.Script.ScriptSource"] = "Inline"
	properties["Octopus.Action.Script.Syntax"] = "PowerShell"
	properties["Octopus.Action.Script.ScriptBody"] = "$customPackageExtractPath = \"#{Octopus.Action.Package[helm_custom_package].ExtractedPath}\"\n$apiGatewayExtractPath = \"#{Octopus.Action.Package[api-gateway-tool].ExtractedPath}\"\n\ncd $apiGatewayExtractPath\n\nvirtualenv apigatewayconfigurationtool \napigatewayconfigurationtool\\Scripts\\activate\n\npip install -r requirements.txt\n\npython api_gateway.py #{Octopus.Environment.Name} $customPackageExtractPath #{kong_admin_url}"

	at := model.ActionTemplate{
		Name:        name,
		Description: description,
		SpaceId:     spaceId,
		ActionType:  "Octopus.Script",
		Packages: []model.ActionTemplatePackages{
			model.ActionTemplatePackages{Name: "helm_custom_package",
				PackageId:           "#{helm_custom_package}",
				FeedId:              feedId,
				AcquisitionLocation: "Server",
				Properties: model.ActionTemplatePackagesProperties{
					Extract: "True",
				},
			},
			model.ActionTemplatePackages{Name: "api-gateway-tool",
				PackageId:           "api-gateway-tool",
				FeedId:              feedId,
				AcquisitionLocation: "Server",
				Properties: model.ActionTemplatePackagesProperties{
					Extract: "True",
				},
			},
		},
		Properties: properties,
	}
	return &at
}
