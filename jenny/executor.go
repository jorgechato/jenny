package jenny

import (
	"fmt"
	"github.com/fatih/color"
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
			SetFilename(len(in) == 3 && (in[2] == "-g" || in[2] == "--global"))
			WriteYaml(jenkins)
		case "login":
			if len(in) == 3 && (in[2] == "-f" || in[2] == "--force-save") {
				WriteYaml(jtmp)
				//TODO: create client if credentials
			}
		case "logout":
			Init(false)
		case "clear":
			RemoveYaml()
			Init(false)
		case "pwd", "user", "use", "project", "uri":
			if len(in) >= 3 {
				third := in[2]
				switch second {
				case "user":
					jtmp.User = third
				case "project":
					jtmp.Project = third
					jenkins.Project = third
				case "pwd":
					jtmp.Password = third
				case "uri":
					jtmp.Uri = third
				}
			}
		}
		return
	case "open":
		if jenkins.IsUri() {
			OpenLink(jenkins)
		}
	default:
		if client == "" {
			color.Yellow("Please login first!")
			return
		} else {
			switch first {
			case "":
			default:
				return
			}
		}
		return
	}
}
