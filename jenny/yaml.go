package jenny

import (
	"fmt"
	"os"
)

func Init(isConfigured bool) {
	if isConfigured {
		fmt.Println("Add a new profile")
	} else {
		fmt.Printf("No %s/.jenny.yml found please type profile.\n", os.Getenv("HOME"))
	}
}

func IsConfigured() bool {
	return false
}
