package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	user, err := user.Current()
	hostname, _ := os.Hostname()

	if err != nil {
		panic(err)
	}
	for {
		prompt := fmt.Sprintf("%s@%s$ ", user.Username, hostname)
		fmt.Print(prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		input = strings.TrimSuffix(input, "\n")

		// Skip an empty input
		if input == "" {
			continue
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}
}

func execInput(input string) error {
	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// add support 'cd' to home with and empty path.
		if len(args) < 2 {
			return errors.New("Path required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// Prepare command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.]
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command return the error.
	return cmd.Run()
}
