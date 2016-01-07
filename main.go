package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var debug = false

// load rules file
func loadRules(ruleFile string) (Rules, error) {
	rulesdata, err := ioutil.ReadFile(ruleFile)

	if err != nil {
		log.Panic("failed to read rules file")
	}
	var rules Rules

	err = yaml.Unmarshal(rulesdata, &rules)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if debug {
		log.Printf("--- rules:\n%v\n\n", rules)
	}

	return rules, err
}

func validFrom(rules Rules, dfile *Dockerfile) bool {
	for _, entry := range rules.From {
		if entry == dfile.From() {
			return true
		}

	}
	return false
}

func isRootUser(rules Rules, dfile *Dockerfile) bool {
	if dfile.User() == "root" || dfile.User() == "" {
		return true
	}
	return false
}

func main() {
	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	dockerfile := os.Getenv("DOCKERFILE")
	ruleFile := os.Getenv("RULESFILE")
	rules, _ := loadRules(ruleFile)
	dfile, err := DockerfileFromPath(dockerfile)
	if err != nil {
		log.Panic(err)
	}
	from := dfile.From()
	log.Println(from)
	if !validFrom(rules, dfile) {
		log.Panic("FROM not valid")
	}
	if isRootUser(rules, dfile) {
		if rules.NotAllowRootUser {
			log.Panic("Running as root")
		}
	}

	os.Exit(0)
}
