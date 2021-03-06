package main

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestValidUser(t *testing.T) {
	dfile, err := DockerfileFromPath("testfiles/Dockerfile.unittest")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.NotNil(t, dfile.User(), "User exists")
	assert.Equal(t, "app", dfile.User(), "User expected")
}

func TestNotUser(t *testing.T) {
	dfile, err := DockerfileFromPath("testfiles/Dockerfile.fail_unittest")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.Equal(t, "", dfile.User(), "User doesn't exists")
}
