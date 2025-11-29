package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jerslf/todo/internal/task"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*task.Tasks, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"add": {
			name:        "add {description}",
			description: "Add new task to the list",
			callback:    commandAdd,
		},
		"list": {
			name:        "list {optional: -a}",
			description: "return a list of all of uncompleted tasks, with option -a to return all tasks regardless of whether or not they are completed",
			callback:    commandList,
		},
		"complete": {
			name:        "complete {id}",
			description: "mark a task as done",
			callback:    commandComplete,
		},
		"delete": {
			name:        "delete {id}",
			description: "remove a task from the list",
			callback:    commandDelete,
		},
		"uncomplete": {
			name:        "uncomplete {id}",
			description: "mark a task as not done",
			callback:    commandUnComplete,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Todo app",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	words[0] = strings.ToLower(words[0])
	return words
}

func StartRepl(ts *task.Tasks) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Todo > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		c, exists := getCommands()[commandName]
		if exists {
			err := c.callback(ts, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
