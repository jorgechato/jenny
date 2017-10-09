package jenny

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

//Executor
func Executor(s string) {
	s = strings.TrimSpace(s)
	in := strings.Split(s, " ")

	first := in[0]
	switch first {
	case "quit", "exit", "q":
		fmt.Println("Bye!")
		os.Exit(0)
		return
	case "help", "h":
		help()
	case "banner", "b":
		Banner()
	case "profile":
		profileActions(in)
		return
	default:
		jenkinsActions(in, first)
	}
}

func profileActions(in []string) {
	typoError(in, 2, func() {
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
			fillProfileActions(in, second)
		}
	})
}

func fillProfileActions(in []string, second string) {
	typoError(in, 3, func() {
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
	})
}

func jenkinsActions(in []string, first string) {
	if client == nil {
		color.Red("Login first!")
		return
	} else {
		typoError(in, 2, func() {
			second := in[1]
			switch first {
			case "build":
				build(client, second)
				//TODO: build
			case "open":
				openLink(client, second)
			case "describe":
				describeJob(client, second)
			case "status":
				typoError(in, 3, func() {
					if isFlag(in, "--last", "-l") {
						getLastStatus(client, second)
					} else {
						getStatus(client, second, stringToInt64(in[2]))
					}
				})
			case "logs":
				typoError(in, 3, func() {
					if isFlag(in, "--last", "-l") {
						getLastLogs(client, second)
					} else {
						getLogs(client, second, stringToInt64(in[2]))
					}
				})
			case "stop":
				typoError(in, 3, func() {
					if isFlag(in, "--last", "-l") {
						stopLastExecution(client, second)
					} else {
						stopExecution(client, second, stringToInt64(in[2]))
					}
				})
			default:
				return
			}
		})
	}
	return
}

func stringToInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return n
	}
	return -1
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

func typoError(in []string, n int, callback func()) {
	if len(in) >= n {
		callback()
	} else {
		color.Red("Wrong command!")
	}
}
