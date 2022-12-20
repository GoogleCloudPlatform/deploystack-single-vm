package main

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/deploystack/dstester"
)

var (
	project  = "ds-test-tf-da43f3d"
	basename = "singlevm"
	debug    = false
)

var Queue = []dstester.GCloudCMD{
	{
		Product: "compute instances",
		Name:    fmt.Sprintf("%s-instances", basename),
		Project: project,
	},
	{
		Product: "compute networks",
		Name:    fmt.Sprintf("%s-network", basename),
		Project: project,
	},
	{
		Product: "compute firewall-rules",
		Name:    "deploystack-allow-ssh",
		Project: project,
	},
}

var vars = map[string]string{
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
}

var tf = dstester.Terraform{
	Dir:  "../terraform",
	Vars: vars,
}

func TestAssertions(t *testing.T) {
	tf.InitApplyForTest(t, debug)
	dstester.TextExistence(Queue, t)
	tf.DestroyForTest(t, debug)
	dstester.TextNonExistence(Queue, t)
}

func TestCreation(t *testing.T) {
	tf.InitApplyForTest(t, debug)
	dstester.TextExistence(Queue, t)
}

func TestDestruction(t *testing.T) {
	tf.DestroyForTest(t, debug)
	dstester.TextNonExistence(Queue, t)
}
