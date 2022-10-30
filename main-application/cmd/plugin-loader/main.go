package main

import (
	"fmt"
	"log"
	"path/filepath"
	"plugin"
	"sync"
	"time"

	"jop-software.de/plugin-test/main-application/pkg/wrapper"
)

func main() {
	plugins, err := filepath.Glob("./plugins/**/*.so")
	if err != nil {
		log.Panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(plugins))

	log.Printf("Loaded %d plugins", len(plugins))

	for _, filename := range plugins {
		p, err := plugin.Open(filename)
		if err != nil {
			log.Panic(err)
		}

		symbol, err := p.Lookup("New")
		if err != nil {
			log.Panic(err)
		}

		plug, err := symbol.(func() (wrapper.Plugin, error))()
		if err != nil {
			log.Panic(err)
		}

		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			plug.Execute()
		}(&wg)
	}

	wg.Wait()
	fmt.Printf("\n")
	time.Sleep(1 * time.Second)
}
