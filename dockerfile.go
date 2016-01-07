package main

import (
	"bytes"
	"github.com/docker/docker/builder/dockerfile/command"
	"github.com/docker/docker/builder/dockerfile/parser"

	"io"
	"io/ioutil"
	"log"
	"strings"
)

func (d *Dockerfile) User() string {
	var ret string
	for _, node := range d.root.Children {
		if node.Value == command.User {
			return strings.Split(node.Original, " ")[1]
		}
	}
	return ret
}

// DockerfileFromPath reads a Dockerfiler from a oath
func DockerfileFromPath(input string) (*Dockerfile, error) {
	payload, err := ioutil.ReadFile(input)
	if err != nil {
		return nil, err
	}
	if debug {
		log.Println(string(payload))
	}
	return DockerfileRead(bytes.NewReader(payload))
}

// DockerfileRead reads a Dockerfile as io.Reader
func DockerfileRead(input io.Reader) (*Dockerfile, error) {
	dockerfile := Dockerfile{}

	root, err := parser.Parse(input)
	if err != nil {
		return nil, err
	}
	dockerfile.root = root

	return &dockerfile, nil
}

// GetFrom returns the current FROM
func (d *Dockerfile) From() string {
	for _, node := range d.root.Children {
		if node.Value == command.From {
			from := strings.Split(node.Original, " ")[1]
			if debug {
				log.Println(from)
			}
			return from
		}
	}

	return ""
}

// String returns a docker-readable Dockerfile
func (d *Dockerfile) String() string {
	lines := []string{}
	for _, child := range d.root.Children {
		lines = append(lines, child.Original)
	}
	return strings.Join(lines, "\n")
}
