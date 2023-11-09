package main

import (
	"crud/initialise"
	"crud/pkg/routes"
)

func init() {
	initialise.LoadEnvVariable()
	initialise.Init()
	routes.Init()
}

func main() {
	routes.RegisterRoutes()
	routes.Serve()
}
