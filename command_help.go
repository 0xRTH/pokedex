package main

import (
	"fmt"
	"os"
)

func exit(config *config, args ...string) error {
	fmt.Println("Bye !")
	os.Exit(0)
	return nil
}
