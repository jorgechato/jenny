package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/jorgechato/jenny/jenny"
)

func main() {
	jenny.Banner()
	//defer func() {
	//fmt.Println("Bye!", recover()) // never go here
	//}()
	defer fmt.Println("Bye!")

	jenny.Init(true)

	menu()
}

func menu() {
	p := prompt.New(
		jenny.Executor,
		jenny.Completer,
		prompt.OptionTitle("Jenny: interactive Jenkins CLI"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionPrefixTextColor(prompt.DarkGray),
		//prompt.OptionPreviewSuggestionTextColor(prompt.DarkGreen),

		prompt.OptionSuggestionBGColor(prompt.Black),
		prompt.OptionDescriptionBGColor(prompt.Black),
		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionDescriptionTextColor(prompt.White),

		prompt.OptionSelectedSuggestionBGColor(prompt.DarkGray),
		prompt.OptionSelectedDescriptionBGColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedDescriptionTextColor(prompt.White),
	)
	p.Run()
}
