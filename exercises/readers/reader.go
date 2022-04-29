//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	cmds := make(map[string]int)

	for {
		input, inputErr := r.ReadString('\n')
		cmd := strings.TrimSuffix(input, "\n")
		if cmd == "" {
			continue
		}

		if strings.ToUpper(cmd) == "Q" {
			break
		}

		switch cmd {
		case "hello":
			fmt.Println("Hello Mr President!")
		case "bye":
			fmt.Println("Goodbyes are sad!")
		}

		cmds[cmd] += 1

		if inputErr != nil {
			fmt.Println("Error reading stdin", inputErr)
		}
	}

	for k, v := range cmds {
		fmt.Printf("Command %q was entered %v time\n", k, v)
	}
}
