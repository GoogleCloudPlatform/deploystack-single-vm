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
	{
		Product: "compute networks",
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
	Dir:  "/Users/tpryan/google/Projects/appinabox/single-vm/terraform",
	Vars: vars,
}

func TestAssertions(t *testing.T) {
	// out, err := tf.Init()
	// if err != nil {
	// 	t.Fatalf("expected no error, got: '%v'", err)
	// }

	// fmt.Printf("out: %s\n", out)

	// out2, err := tf.Exec(vars)
	// if err != nil {
	// 	t.Fatalf("expected no error, got: '%v'", err)
	// }

	// fmt.Printf("out2: %s\n", out2)

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

	// testsNotExists := map[string]struct {
	// 	input dstester.GCloudCMD
	// }{}
	// for _, v := range Queue {
	// 	testsNotExists[fmt.Sprintf("Test %s does not exist", v.Name)] = struct {
	// 		input dstester.GCloudCMD
	// 	}{v}
	// }

	// for name, tc := range testsNotExists {
	// 	t.Run(name, func(t *testing.T) {
	// 		_, err := tc.input.Describe()
	// 		if err == nil {
	// 			t.Fatalf("expected error, got no error")
	// 		}
	// 	})
	// }
}
