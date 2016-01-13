package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var debug = false
var rules Rules

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
	lines := []string{}
	valid := true
	if !v.validFrom() {
		valid = false
		lines = append(lines, "FROM not valid")
	}

	if v.isRootUser() {
		if v.Rules.NotAllowRootUser {
			valid = false
			lines = append(lines, "Running as root")
		}
	}

	return valid, strings.Join(lines, "\n")

}

func main() {

	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	rules, _ = loadRulesFromFile("rules.yaml")
	log.Println("listening in port 8080")
	http.HandleFunc("/validate", validateHandler)
	http.HandleFunc("/setup", uploadHandler)
	http.HandleFunc("/rules", uploadRulesHandler)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
