package cmd

import "os"

var octopusHost, octopusApiKey string

func init() {
	octopusHost = os.Getenv("OCTOPUS_HOST")
	if octopusHost == "" {
		panic("OCTOPUS_HOST not found in environment. Please set it as a environment variable.")
	}
	octopusApiKey = os.Getenv("OCTOPUS_APIKEY")
	if octopusApiKey == "" {
		panic("OCTOPUS_APIKEY not found in environment. Please set it as a environment variable.")
	}
}
