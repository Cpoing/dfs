package main

import (
	"fmt"
)

func (cli *DistributedFileStorageCLI) initSystem(args []string) {
	fmt.Println("Initializing Distributed File Storage System...")
	fmt.Println("- Generating configuration")
	fmt.Println("- Setting up storage nodes")
	fmt.Println("- Configuring encryption")
	fmt.Println("Initialization complete!")
}

func (cli *DistributedFileStorageCLI) exitApplication(args []string) {
	fmt.Println("Exiting Distributed File Storage System. Goodbye!")
	cli.running = false
}

func (cli *DistributedFileStorageCLI) showHelp(args []string) {
	fmt.Println("Available commands: ")
	for name, cmd := range cli.commands {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}
}
