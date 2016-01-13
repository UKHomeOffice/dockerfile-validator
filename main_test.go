package main

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestValidFrom(t *testing.T) {
	rules, _ := loadRulesFromFile("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile")
	v := Validation{rules, dfile}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.True(t, v.validFrom(), "FROM entry is valid")
}

func TestDockerfileRead(t *testing.T) {

	dfile, err := DockerfileFromPath("Dockerfile")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.NotNil(t, dfile, "Read Dockerfile")
}

func TestFailFrom(t *testing.T) {
	rules, _ := loadRulesFromFile("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile.fail_unittest")
	v := Validation{rules, dfile}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.False(t, v.validFrom(), "FROM entry is valid")
}

func TestIsRootUser(t *testing.T) {
	rules, _ := loadRulesFromFile("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile.fail_unittest")
	v := Validation{rules, dfile}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.True(t, v.isRootUser(), "Runs as root")
}

func TestIsNOTRootUser(t *testing.T) {
	rules, _ := loadRulesFromFile("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile.unittest")
	v := Validation{rules, dfile}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.False(t, v.isRootUser(), "Runs as not root")
}

func TestValidate(t *testing.T) {
	rules, _ := loadRulesFromFile("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile.unittest")
	v := Validation{rules, dfile}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r, _ := v.validate()
	assert.True(t, r, "Dockerfile is valid")
}
