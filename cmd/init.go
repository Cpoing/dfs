package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
)

type Command struct {
  Name string
  Description string
  Handler func([]string)
}

type DistributedFileStorageCLI struct {
  commands map[string]Command
  running bool
}

func NewDistributedFileStorageCLI() *DistributedFileStorageCLI{
  cli := &DistributedFileStorageCLI{
    commands: make(map[string]Command),
    running: true,
  }

  cli.RegisterCommand(Command{
    Name: "init",
    Description: "Initializing the distributed file storage system",
    Handler: cli.initSystem,
  })

  cli.RegisterCommand(Command{
    Name: "exit",
    Description: "Exit the application",
    Handler: cli.exitApplication,
  })

  return cli
}

func (cli *DistributedFileStorageCLI) RegisterCommand(cmd Command) {
  cli.commands[cmd.Name] = cmd
}

func (cli *DistributedFileStorageCLI) ProcessCommand(input string) {
  parts := strings.Fields(input)

  cmdName := parts[0]

  cmd, exists := cli.commands[cmdName]
  if !exists {
    fmt.Printf("Unknown command: %s. Type 'help' for available commands.\n", cmdName)
    return
  }

  cmd.Handler(parts[1:])
}

func (cli *DistributedFileStorageCLI) Start() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Distrubted File Storage System CLI")
  fmt.Println("Type 'help' for available commands")

  for cli.running {
    fmt.Print("dfs> ")

    input, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("Error reading input", err)
      continue
    }

    input = strings.TrimSpace(input)

    cli.ProcessCommand(input)
  }
}
