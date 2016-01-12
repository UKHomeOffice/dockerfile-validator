package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"

	"os"
	"strconv"
)

var debug = false

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
	rulesdata, _ := ioutil.ReadFile(ruleFile)
	return loadRules(rulesdata)
}

func (v Validation) validFrom() bool {
	for _, entry := range v.Rules.From {
		if entry == v.Dockerfile.From() {
			return true
		}
	}
	return false
}

func (v Validation) isRootUser() bool {
	if v.Dockerfile.User() == "root" || v.Dockerfile.User() == "" {
		return true
	}
	return false
}

func (v Validation) validate() (bool, string) {

	from := v.Dockerfile.From()
	log.Println(from)
	if !v.validFrom() {
		return false, "FROM not valid"
	}

	if v.isRootUser() {
		if v.Rules.NotAllowRootUser {
			return false, "Running as root"
		}
	}
	return true, ""

}

func main() {
	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	log.Println("listening in port 8080")
	http.HandleFunc("/validate", uploadHandler)
	http.ListenAndServe(":8080", nil)
}
