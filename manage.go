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
		//kube.Executor,
		//kube.Completer,
		prompt.OptionTitle("jenny: interactive Jenkins CLI"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
