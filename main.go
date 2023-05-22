package main

import "github.com/Reinhadt/celeritas"

// holds all our app needs
type application struct {
	App *celeritas.Celeritas
}

func main() {
	initApplication()
}
