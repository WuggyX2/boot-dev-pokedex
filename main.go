package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	cmds := registerCmds()

	for {
		fmt.Print("Pokdedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		userCommand := input[0]

		command, exists := cmds[userCommand]

		if exists {
			err := command.callback()

			if err != nil {
				fmt.Println("Command execution failed")
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
