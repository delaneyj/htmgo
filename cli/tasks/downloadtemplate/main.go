package downloadtemplate

import (
	"flag"
	"fmt"
	"github.com/maddalax/htmgo/cli/tasks/process"
	"github.com/maddalax/htmgo/cli/tasks/run"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func deleteAllExceptTemplate(outPath string, excludeDir string) {
	// List all files and directories in the root folder
	files, err := os.ReadDir(outPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Iterate through each item in the root folder
	for _, file := range files {
		// Skip the excluded directory
		if file.Name() == excludeDir {
			continue
		}

		// Get full path
		fullPath := filepath.Join(outPath, file.Name())

		err := os.RemoveAll(fullPath)
		if err != nil {
			fmt.Printf("Error removing %s: %v\n", fullPath, err)
		}
	}
}

func DownloadTemplate(outPath string) {
	cwd, _ := os.Getwd()

	flag.Parse()

	outPath = strings.ReplaceAll(outPath, "\n", "")
	outPath = strings.ReplaceAll(outPath, " ", "-")
	outPath = strings.ToLower(outPath)

	if outPath == "" {
		fmt.Println("Please provide a name for your app.")
		return
	}

	excludeDir := "starter-template"

	install := exec.Command("git", "clone", "https://github.com/maddalax/htmgo", "--depth=1", outPath)
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	err := install.Run()

	if err != nil {
		println("Error downloading template %v\n", err)
		return
	}

	deleteAllExceptTemplate(outPath, excludeDir)

	newDir := filepath.Join(cwd, outPath)

	commands := [][]string{
		{"cp", "-vaR", fmt.Sprintf("%s/.", excludeDir), "."},
		{"rm", "-rf", excludeDir},
		{"go", "get", "github.com/maddalax/htmgo/framework"},
		{"go", "get", "github.com/maddalax/htmgo/framework-ui"},
		{"git", "init"},
	}

	for _, command := range commands {
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Dir = newDir
		cmd.Stdout = nil
		cmd.Stderr = nil
		err = cmd.Run()
		if err != nil {
			println("Error executing command %s\n", err.Error())
			return
		}
	}

	fmt.Printf("Setting up the project in %s\n", newDir)
	process.SetWorkingDir(newDir)
	run.Setup()
	process.SetWorkingDir("")

	fmt.Println("Template downloaded successfully.")
	fmt.Println("To start the development server, run the following commands:")
	fmt.Printf("cd %s && htmgo watch\n", outPath)

	fmt.Printf("To build the project, run the following command:\n")
	fmt.Printf("cd %s && htmgo build\n", outPath)
}
