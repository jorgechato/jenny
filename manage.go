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

func main() {
	version = "a0.0"
	revision = "alpha"

	fmt.Println(jenny.Banner())
	fmt.Printf("Jenny %s (rev-%s) powered with ❤ by Jorge Chato\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program..")
	defer fmt.Println("Bye!")

	Init()

	Menu()
}

func Init() {
	isConfigured := jenny.IsConfigured()

	if !isConfigured {
		jenny.Init(isConfigured)
	}

	//j := jenny.Jenkins{
	//Uri:      "uri",
	//User:     "user",
	//Password: "sd",
	//}
	//correct := j.PasswordMatch("sd")
}

func Menu() {
	p := prompt.New(
		jenny.Executor,
		jenny.Completer,
		prompt.OptionTitle("Jenny: interactive Jenkins CLI"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionPrefixTextColor(prompt.Red),
		prompt.OptionPreviewSuggestionTextColor(prompt.DarkGreen),

		prompt.OptionSuggestionBGColor(prompt.DefaultColor),
		prompt.OptionDescriptionBGColor(prompt.DefaultColor),
		prompt.OptionSuggestionTextColor(prompt.DefaultColor),
		prompt.OptionDescriptionTextColor(prompt.DefaultColor),

		prompt.OptionSelectedSuggestionBGColor(prompt.DarkGreen),
		prompt.OptionSelectedDescriptionBGColor(prompt.DefaultColor),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGreen),
	)
	p.Run()
}
