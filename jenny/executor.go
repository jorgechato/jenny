package jenny

import (
	"fmt"
	"os"
	"strings"
	//"gopkg.in/yaml.v2"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	fmt.Println(s)

	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	} else if s == "profile" {
		//TODO: isConfigure discard config, if not exit jenny
		if !IsConfigured() {
			os.Exit(0)
		} else {
			//TODO: relunch default config
		}
		return
	} else if s == "save" {
		//TODO: save yaml
		return
	}
}
