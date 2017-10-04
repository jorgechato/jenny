package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/jorgechato/jenny/jenny"
)

var (
	version  string
	revision string
)

// executor executes command and print the output.
func executor(in string) {
	fmt.Println("Your input: " + in)
}

// completer returns the completion items from user input.
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "user table"},
		{Text: "sites", Description: "sites table"},
		{Text: "articles", Description: "articles table"},
		{Text: "comments", Description: "comments table"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Printf("Jenny %s (rev-%s)\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program..")
	defer fmt.Println("Bye!")

	j := jenny.Jenkins{
		Uri:      "uri",
		User:     "user",
		Password: "sd",
	}

	fmt.Println(j.PasswordMatch("sd"))
	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("jenny: interactive Jenkins CLI"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
