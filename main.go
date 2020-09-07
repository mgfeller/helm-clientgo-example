// Example to demonstrate helm chart installation using helm client-go
// Most of the code is copied from https://github.com/helm/helm repo

package main

import (
	"github.com/PrasadG193/helm-clientgo-example/helm"
	"helm.sh/helm/v3/pkg/cli"
)

var settings *cli.EnvSettings

func main() {
	consulChart := helm.ChartSpec{
		Url:         "https://helm.releases.hashicorp.com",
		RepoName:    "hashicorp",
		ChartName:   "consul",
		ReleaseName: "hashicorp",
		Namespace:   "consul-test-helm-3",
		Args: map[string]string{
			// comma separated values to set
			"set": "server.replicas=1,server.bootstrapExpect=1",
		},
	}
	settings = helm.Settings(consulChart.Namespace)
	// Add helm repo
	helm.RepoAdd(consulChart.RepoName, consulChart.Url, settings)
	// Update charts from the helm repo
	helm.RepoUpdate(settings)
	// Install charts
	helm.InstallChart(consulChart.ReleaseName, consulChart.RepoName, consulChart.ChartName, consulChart.Args, settings)
}
