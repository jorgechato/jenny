package jenny

import "fmt"

func help() {
	fmt.Println("Usage:\n")
	for _, action := range actions {
		fmt.Printf(" %s\t\t%s\n", action.Text, action.Description)
	}

	fmt.Println("\nMain modes of profile:\n")
	for _, action := range profile {
		fmt.Printf(" %s\t\t%s\n", action.Text, action.Description)
	}

	fmt.Println("\nProfile flags:\n")
	for _, action := range flagInit {
		fmt.Printf(" %s\t\t%s\n", action.Text, action.Description)
	}
	for _, action := range flagLogout {
		fmt.Printf(" %s\t\t%s\n", action.Text, action.Description)
	}
	for _, action := range flagShow {
		fmt.Printf(" %s\t\t%s\n", action.Text, action.Description)
	}

	fmt.Println("\nJenkins operators flags:\n")
	for _, action := range flagJenkins {
		fmt.Printf(" %s\t\t%s\n", action.Text, action.Description)
	}
}
