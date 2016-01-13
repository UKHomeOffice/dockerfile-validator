package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestString(t *testing.T) {
	rulesFile, err := loadRulesFromFile("rules.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.NotNil(t, rulesFile, "Read and validate rules file")
}

func TestLoadRules(t *testing.T) {
	rulesFile, err := loadRulesFromFile("rules.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.NotNil(t, rulesFile, "Read and validate rules file")
}

func TestUnmarshal(t *testing.T) {
	rules, _ := loadRulesFromFile("rules.yaml")
	assert.True(t, rules.RootUser, "Unmashalling boolean value with mapping")
}
