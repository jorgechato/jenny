package jenny

import (
	"fmt"
	"os"
	"strings"
	//"gopkg.in/yaml.v2"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	in := strings.Split(s, " ")

	first := in[0]
	switch first {
	case "quit", "exit", "q":
		fmt.Println("Bye!")
		os.Exit(0)
		return
	case "profile":
		second := in[1]
		switch second {
		case "cancel":
			if !IsConfigured() {
				os.Exit(0)
				return
			} else {
				//TODO: lunch default config
			}
		case "save":
			if len(in) == 3 && (in[2] == "-f" || in[2] == "--force-save") {
				//TODO: save yaml
			}
		case "pwd", "user", "use", "name", "uri":
			//third := in[2]
			//TODO: set profile
		}
		return
	//TODO: add jenkins api and credential middleware
	case "":
		return
	default:
		return
	}
}
