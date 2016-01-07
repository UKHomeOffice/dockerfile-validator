package main

import (
	"github.com/docker/docker/builder/dockerfile/command"
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
