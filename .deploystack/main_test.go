package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/deploystack/dstester"
)

var project = "ds-test-tf-da43f3d"

var Queue = []dstester.GCloudCMD{
	{
		Product: "projects",
		Name:    "ds-test-tf-da43f3d",
		Project: project,
	},
	{
		Product: "compute instances",
		Name:    "test-instance",
		Project: project,
	},
}

var vars = map[string]string{
	"REGION":       "us-central1",
	"ZONE":         "us-central1-a",
	"BASENAME":     "singlevm",
	"DISKSIZE":     "200",
	"DISKTYPE":     "pd-standard",
	"DISKIMAGE":    "debian-cloud/debian-11-bullseye-v20220519",
	"MACHINETYPE":  "n1-standard-1",
	"NAME":         "singlevm-instance",
	"TAGS":         "[\"http-server\",\"https-server\"]",
	"terraformDIR": "terraform",
}

var tf = dstester.Terraform{
	Dir:  "../terraform",
	Vars: vars,
}

func TestAssertions(t *testing.T) {
	out, err := tf.Init()
	if err != nil {
		t.Fatalf("expected no error, got: '%v'", err)
	}

	fmt.Printf("out: %s\n", out)

	testsExists := map[string]struct {
		input dstester.GCloudCMD
		want  string
	}{}
	for _, v := range Queue {
		testsExists[fmt.Sprintf("Test %s exists", v.Name)] = struct {
			input dstester.GCloudCMD
			want  string
		}{v, v.Name}
	}

	for name, tc := range testsExists {
		t.Run(name, func(t *testing.T) {
			got, err := tc.input.Describe()
			if err != nil {
				t.Fatalf("expected no error, got: '%v'", err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: '%v', got: '%v'", tc.want, got)
			}
		})
	}

	testsNotExists := map[string]struct {
		input dstester.GCloudCMD
	}{}
	for _, v := range Queue {
		testsNotExists[fmt.Sprintf("Test %s does not exist", v.Name)] = struct {
			input dstester.GCloudCMD
		}{v}
	}

	for name, tc := range testsNotExists {
		t.Run(name, func(t *testing.T) {
			_, err := tc.input.Describe()
			if err == nil {
				t.Fatalf("expected error, got no error")
			}
		})
	}
}
