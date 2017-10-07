package jenny

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
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
			Init(false)
		case "show":
			uncover := isFlag(in, "--uncover", "-u")
			printJenkins(jenkins, uncover)
		case "save":
			force, changeFilename := areSaveFlags(in)
			setFilename(changeFilename)
			writeYaml(jenkins, force)
		case "login":
			jenkins.login()
		case "logout":
			Init(false)
			if isFlag(in, "--force", "-f") {
				writeYaml(jenkins, false)
			}
		case "clear":
			removeYaml()
			Init(false)
		case "pwd", "user", "project", "uri":
			if len(in) >= 3 {
				third := in[2]
				switch second {
				case "user":
					jenkins.User = third
				case "project":
					jenkins.Project = third
				case "pwd":
					jenkins.Password = third
				case "uri":
					jenkins.Uri = third
				}
			}
		}
		return
	case "open":
		if jenkins.IsUri() {
			openLink(jenkins)
		}
	default:
		if client == nil {
			color.Red("Login first!")
			return
		} else {
			if len(in) >= 2 {
				second := in[1]
				switch first {
				case "status":
					if isFlag(in, "--last", "-l") {
						getLastStatus(client, second)
					} else {
						if len(in) >= 3 {
							getStatus(client, second, stringToInt64(in[2]))
						}
					}
				case "logs":
					if isFlag(in, "--last", "-l") {
						getLastLogs(client, second)
					} else {
						if len(in) >= 3 {
							getLogs(client, second, stringToInt64(in[2]))
						}
					}
				case "stop":
					if isFlag(in, "--last", "-l") {
						stopLastExecution(client, second)
					} else {
						if len(in) >= 3 {
							stopExecution(client, second, stringToInt64(in[2]))
						}
					}
				default:
					return
				}
			}
		}
		return
	}
}

func stringToInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return n
	}
	return 1
}

func isFlag(in []string, s string, a string) bool {
	isThere := false
	for _, action := range in {
		isThere = action == a || action == s
		if isThere {
			break
		}
	}
	return isThere
}

func areSaveFlags(in []string) (bool, bool) {
	return isFlag(in, "--force", "-f"), isFlag(in, "--global", "-g")
}
