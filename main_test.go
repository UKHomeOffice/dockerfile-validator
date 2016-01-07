package main

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadRules(t *testing.T) {
	rulesFile, err := loadRules("rules.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.NotNil(t, rulesFile, "Read and validate rules file")

}

func TestValidFrom(t *testing.T) {
	rules, _ := loadRules("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.True(t, validFrom(rules, dfile), "FROM entry is valid")
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
	rules, _ := loadRules("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile.fail_unittest")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.False(t, validFrom(rules, dfile), "FROM entry is valid")
}

func TestIsRootUser(t *testing.T) {
	rules, _ := loadRules("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile.fail_unittest")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.True(t, isRootUser(rules, dfile), "Runs as root")
}

func TestIsNOTRootUser(t *testing.T) {
	rules, _ := loadRules("rules.yaml")
	dfile, err := DockerfileFromPath("Dockerfile.unittest")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.False(t, isRootUser(rules, dfile), "Runs as not root")
}
