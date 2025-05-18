package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	next     string
	previous string
}

func cleanInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	nonEmpty := []string{}

	for value := range strings.SplitSeq(trimmed, " ") {
		if len(value) > 0 {
			lowered := strings.ToLower(value)
			nonEmpty = append(nonEmpty, lowered)
		}
	}

	return nonEmpty
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}
	cmds := registerCmds(&cfg)

	for {
		fmt.Print("Pokdedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		userCommand := input[0]

		command, exists := cmds[userCommand]

		if exists {
			err := command.callback(input[1:])

			if err != nil {
				fmt.Println("Command execution failed")
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
