package main

import "fmt"

func help(config *config, args ...string) error {
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("%s : %s\n", command.name, command.description)
	}
	return nil
}
