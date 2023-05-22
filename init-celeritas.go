package main

import (
	"log"
	"os"

	"github.com/Reinhadt/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd() // get working dir
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)

	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"
	cel.Debug = true
	app := &application{
		App: cel,
	}

	return app

}