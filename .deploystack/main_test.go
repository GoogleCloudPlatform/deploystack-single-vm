package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/deploystack"
	"github.com/GoogleCloudPlatform/deploystack/dstester"
)

var (
	ops        = dstester.NewOperationsSet()
	project, _ = deploystack.ProjectID()
	basename   = "singlevm"
	debug      = false

	tf = dstester.Terraform{
		Dir: "../terraform",
		Vars: map[string]string{
			"project_id":            project,
			"project_number":        "753592922120",
			"region":                "us-central1",
			"zone":                  "us-central1-a",
			"basename":              "singlevm",
			"instance-disksize":     "200",
			"instance-disktype":     "pd-standard",
			"instance-image":        "debian-cloud/debian-11-bullseye-v20220519",
			"instance-machine-type": "n1-standard-1",
			"instance-name":         "singlevm-instance",
			"instance-tags":         "[\"http-server\",\"https-server\"]",
		},
	}

	resources = dstester.Resources{
		Project: project,
		Items: []dstester.Resource{
			{
				Product: "compute instances",
				Name:    fmt.Sprintf("%s-instance", basename),
			},
			{
				Product: "compute networks",
				Name:    fmt.Sprintf("%s-network", basename),
			},
			{
				Product: "compute firewall-rules",
				Name:    "deploystack-allow-ssh",
			},
		},
	}
)

func init() {
	if os.Getenv("debug") != "" {
		debug = true
	}
}

func TestListCommands(t *testing.T) {
	resources.Init()
	dstester.DebugCommands(t, tf, resources)
}

func TestStack(t *testing.T) {
	dstester.TestStack(t, tf, resources, ops, debug)
}

func TestClean(t *testing.T) {
	if os.Getenv("clean") == "" {
		t.Skip("Clean must be very intentionally called")
	}

	resources.Init()
	dstester.Clean(t, tf, resources)
}
