package main

import "github.com/ExtraProjects860/Project-Device-Mobile/routes"

func main() {
	router := routes.InitializeRouter()

	router.Run(":5050")
}
