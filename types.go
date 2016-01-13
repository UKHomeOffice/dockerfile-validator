package main

import (
	"github.com/docker/docker/builder/dockerfile/parser"
)

type Dockerfile struct {
	root *parser.Node
}

type Rules struct {
	From     []string
	RootUser bool `yaml:"RootUser"`
}

type Validation struct {
	Rules      Rules
	Dockerfile *Dockerfile
}
