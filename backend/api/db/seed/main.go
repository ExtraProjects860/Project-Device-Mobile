package main

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
)

func main() {
	logger := config.NewLogger("seed")

	_, db, err := config.Init()
	if err != nil {
		panic(fmt.Errorf("failed to init config: %v", err))
	}

	err = config.ResetDB(db, logger)
	if err != nil {
		panic(fmt.Errorf("failed to reset database: %v", err))
	}
	seeds(db, logger)
}
