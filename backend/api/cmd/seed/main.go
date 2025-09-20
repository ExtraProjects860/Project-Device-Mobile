package main

import "github.com/ExtraProjects860/Project-Device-Mobile/seed"

func main() {
	seed.InitializeHandler()
	seed.ResetDB()
	seed.Seeds()
}
