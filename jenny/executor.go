package jenny

import (
	"fmt"
	"os"
	"strings"
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
			if jenkins.IsEmpty() {
				fmt.Println("Configuration aborted")
				fmt.Println("Bye!")
				os.Exit(0)
				return
			}
			jtmp = jenkins
		case "show":
			uncover := len(in) == 3 && (in[2] == "-u" || in[2] == "--uncover")
			PrintJenkins(jenkins, uncover)
		case "save":
			jenkins = jtmp
			if len(in) == 3 && (in[2] == "-f" || in[2] == "--force-save") {
				WriteYaml(jtmp)
			}
		case "pwd", "user", "use", "name", "uri":
			if len(in) >= 3 {
				third := in[2]
				switch second {
				case "user":
					jtmp.User = third
				case "name":
					jtmp.Name = third
				case "pwd":
					jtmp.Password = third
				case "uri":
					jtmp.Uri = third
				case "use":
					//TODO: load new profile
				}
			}
		}
		return
	default:
		if jenkins.IsEmpty() {
			fmt.Println("You don't have the right credentials (fill the profile).")
			return
		}
		switch first {
		case "":
		default:
			return
		}
		return
	}
}
