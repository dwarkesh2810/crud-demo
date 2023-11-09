package main

import (
	"crud/pkg/config"
	"crud/pkg/routes"
)

func init() {
	config.LoadEnvVariable()
	config.Init()
	routes.Init()
}

func main() {
	routes.RegisterRoutes()
	routes.Serve()
}
