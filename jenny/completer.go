package jenny

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	return argumentsCompleter(args)
}

var commands = []prompt.Suggest{
	{Text: "get", Description: "Display one or many resources"},
	{Text: "describe", Description: "Show details of a specific resource or group of resources"},
	{Text: "create", Description: "Create a resource by filename or stdin"},
	{Text: "delete", Description: "Delete resources by filenames, stdin, resources and names, or by resources and label selector."},
	{Text: "edit", Description: "Edit a resource on the server"},
	{Text: "apply", Description: "Apply a configuration to a resource by filename or stdin"},
	{Text: "logs", Description: "Print the logs for a job or an event."},
	{Text: "run", Description: "Run a particular job."},
	{Text: "version", Description: "Print the Jenkins version information."},
	{Text: "explain", Description: "Documentation of resources."},
	{Text: "exit", Description: "Exit this program"},
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}
	return []prompt.Suggest{}
}
