package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed Taskfile.yml
var taskFile string

func main() {
	commandMap := make(map[string]*flag.FlagSet)
	commands := []string{"template", "run", "build", "setup"}

	for _, command := range commands {
		commandMap[command] = flag.NewFlagSet(command, flag.ExitOnError)
	}

	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("Usage: htmgo [%s]", strings.Join(commands, " | ")))
		os.Exit(1)
	}

	err := commandMap[os.Args[1]].Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	// Install the latest version of Task
	install := exec.Command("go", "install", "github.com/go-task/task/v3/cmd/task@latest")

	err = install.Run()
	if err != nil {
		fmt.Printf("Error installing task: %v\n", err)
		return
	}

	temp, err := os.CreateTemp("", "Taskfile.yml")

	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		return
	}

	os.WriteFile(temp.Name(), []byte(taskFile), 0644)

	// Define the command and arguments
	cmd := exec.Command("task", "-t", temp.Name(), os.Args[1])
	// Set the standard output and error to be the same as the Go program
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running task command: %v\n", err)
		return
	}
}
