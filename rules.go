package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// load rules file
func loadRules(rulesdata []byte) (Rules, error) {
	var rules Rules
	err := yaml.Unmarshal(rulesdata, &rules)

	if debug {
		log.Printf("--- rules:\n%v\n\n", rules)
	}

	return rules, err
}

// load rules file
func loadRulesFromFile(ruleFile string) (Rules, error) {
	var rules Rules
	if _, err := os.Stat(ruleFile); err == nil {
		rulesdata, _ := ioutil.ReadFile(ruleFile)
		return loadRules(rulesdata)
	} else {
		return rules, err
	}
}

// String returns a docker-readable Dockerfile
func (d *Rules) String() string {
	lines := []string{}
	for _, child := range d.From {
		lines = append(lines, "FROM "+child)
	}
	user := fmt.Sprintf("Root user allowed: %v", d.NotAllowRootUser)
	lines = append(lines, user)
	return strings.Join(lines, "\n")
}
