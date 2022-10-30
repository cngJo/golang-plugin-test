package main

import (
	"fmt"

	"jop-software.de/plugin-test/main-application/pkg/wrapper"
	"jop-software.de/plugin/first/internal/calc"
)

const ModuleName = "First"

type Module struct{}

func New() (wrapper.Plugin, error) {
	return Module{}, nil
}

// Execute holds the logic for the plugin
func (s Module) Execute() {
	fmt.Printf("%d + %d = %d\n", 10, 20, calc.Add(10, 20))
}
