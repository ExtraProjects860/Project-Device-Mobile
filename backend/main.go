package main

import "github.com/ExtraProjects860/Project-Device-Mobile/router"

func main() {
	router := router.InitializeRouter()

	router.Run(":5050")
}
