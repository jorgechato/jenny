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

	w := d.GetWordBeforeCursor()
	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, strings.HasPrefix(w, "--"))
	}

	return argumentsCompleter(excludeOptions(args))
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(actions, args[0], true)
	}

	first := args[0]
	if len(args) == 2 {
		second := args[1]
		switch first {
		case "profile":
			return prompt.FilterHasPrefix(profile, second, true)
		case "status", "logs", "stop":
			return prompt.FilterHasPrefix(jobsNames, second, true)
		}
	}

	return []prompt.Suggest{}
}
