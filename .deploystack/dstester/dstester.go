package dstester

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/GoogleCloudPlatform/deploystack"
)

func SectionOpen(desc string) {
	fmt.Println(deploystack.Divider)
	fmt.Printf("%s%s%s\n", deploystack.TERMCYAN, desc, deploystack.TERMCLEAR)
	fmt.Println(deploystack.Divider)
}

type Terraform struct {
	Dir  string
	Vars map[string]string
}

func (t Terraform) Init() (string, error) {
	cmd := exec.Command("terraform")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Args = append(cmd.Args, fmt.Sprintf("-chdir=%s", t.Dir))
	cmd.Args = append(cmd.Args, "init")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("stderr %v\n", stderr.String())
		fmt.Printf("stdout %v\n", stdout.String())
		return "", fmt.Errorf("error: %s \nextra: '%s'", err, stdout.String())
	}
	out := strings.TrimSpace(stdout.String())

	return out, nil
}

func (t Terraform) Exec(vars map[string]string) (string, error) {
	cmd := exec.Command("terraform")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Args = append(cmd.Args, fmt.Sprintf("-chdir=%s", t.Dir))
	cmd.Args = append(cmd.Args, "apply")
	cmd.Args = append(cmd.Args, "-auto-approve")

	for i, v := range vars {
		cmd.Args = append(cmd.Args, "-var")
		cmd.Args = append(cmd.Args, fmt.Sprintf("%s=%s", i, v))
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd %v\n", cmd.String())
		fmt.Printf("stderr %v\n", stderr.String())
		fmt.Printf("stdout %v\n", stdout.String())
		return "", fmt.Errorf("error: %s \nextra: '%s'", err, stdout.String())
	}
	out := strings.TrimSpace(stdout.String())

	return out, nil
}

type GCloudCMD struct {
	Product string
	Name    string
	Field   string
	Append  string
	Project string
}

func (g *GCloudCMD) Describe() (string, error) {
	cmd := exec.Command("gcloud")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	for _, v := range strings.Split(g.Product, " ") {
		cmd.Args = append(cmd.Args, v)
	}

	cmd.Args = append(cmd.Args, "describe", g.Name)

	if len(g.Append) > 0 {
		for _, v := range strings.Split(g.Append, " ") {
			cmd.Args = append(cmd.Args, v)
		}
	}

	if g.Field == "" {
		g.Field = "name"
	}

	cmd.Args = append(cmd.Args, fmt.Sprintf("--format=value(%s)", g.Field))

	if g.Project != "" {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--project=%s", g.Project))
	}

	fmt.Printf("cmd %v\n", cmd.String())

	dat, err := cmd.Output()
	if err != nil {
		fmt.Printf("stderr %v\n", stderr.String())
		fmt.Printf("stdout %v\n", stdout.String())
		return "", fmt.Errorf("error: %s \nextra: '%s'", err, string(dat))
	}
	out := strings.TrimSpace(string(dat))

	return out, nil
}
