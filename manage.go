package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/jorgechato/jenny/jenny"
	"math/rand"
	"time"
)

var (
	version  string
	revision string
)

var banner = []string{`
    ___       ___       ___       ___       ___   
   /\  \     /\  \     /\__\     /\__\     /\__\  
  _\:\  \   /::\  \   /:| _|_   /:| _|_   |::L__L 
 /\/::\__\ /::\:\__\ /::|/\__\ /::|/\__\  |:::\__\
 \::/\/__/ \:\:\/  / \/|::/  / \/|::/  /  /:;;/__/
  \/__/     \:\/  /    |:/  /    |:/  /   \/__/   
             \/__/     \/__/     \/__/            
`,
	`
   __     ______     __   __     __   __     __  __    
  /\ \   /\  ___\   /\ "-.\ \   /\ "-.\ \   /\ \_\ \   
 _\_\ \  \ \  __\   \ \ \-.  \  \ \ \-.  \  \ \____ \  
/\_____\  \ \_____\  \ \_\\"\_\  \ \_\\"\_\  \/\_____\ 
\/_____/   \/_____/   \/_/ \/_/   \/_/ \/_/   \/_____/ 
`,
	`
     ██╗███████╗███╗   ██╗███╗   ██╗██╗   ██╗
     ██║██╔════╝████╗  ██║████╗  ██║╚██╗ ██╔╝
     ██║█████╗  ██╔██╗ ██║██╔██╗ ██║ ╚████╔╝ 
██   ██║██╔══╝  ██║╚██╗██║██║╚██╗██║  ╚██╔╝  
╚█████╔╝███████╗██║ ╚████║██║ ╚████║   ██║   
 ╚════╝ ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═══╝   ╚═╝   
`,
	`
::::::'##:'########:'##::: ##:'##::: ##:'##:::'##:
:::::: ##: ##.....:: ###:: ##: ###:: ##:. ##:'##::
:::::: ##: ##::::::: ####: ##: ####: ##::. ####:::
:::::: ##: ######::: ## ## ##: ## ## ##:::. ##::::
'##::: ##: ##...:::: ##. ####: ##. ####:::: ##::::
 ##::: ##: ##::::::: ##:. ###: ##:. ###:::: ##::::
. ######:: ########: ##::. ##: ##::. ##:::: ##::::
:......:::........::..::::..::..::::..:::::..:::::
`,
	`
     @@@  @@@@@@@@  @@@  @@@  @@@  @@@  @@@ @@@  
     @@@  @@@@@@@@  @@@@ @@@  @@@@ @@@  @@@ @@@  
     @@!  @@!       @@!@!@@@  @@!@!@@@  @@! !@@  
     !@!  !@!       !@!!@!@!  !@!!@!@!  !@! @!!  
     !!@  @!!!:!    @!@ !!@!  @!@ !!@!   !@!@!   
     !!!  !!!!!:    !@!  !!!  !@!  !!!    @!!!   
     !!:  !!:       !!:  !!!  !!:  !!!    !!:    
!!:  :!:  :!:       :!:  !:!  :!:  !:!    :!:    
::: : ::   :: ::::   ::   ::   ::   ::     ::    
 : :::    : :: ::   ::    :   ::    :      :     
`}

func main() {
	version = "a0.0"
	revision = "alpha"

	rand.Seed(time.Now().UnixNano())
	fmt.Println(banner[rand.Int()%len(banner)])
	fmt.Printf("Jenny %s (rev-%s) powered with ❤ by Jorge Chato\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program..")
	defer fmt.Println("Bye!")

	p := prompt.New(
		jenny.Executor,
		jenny.Completer,
		prompt.OptionTitle("Jenny: interactive Jenkins CLI"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionPrefixTextColor(prompt.Red),
		prompt.OptionPreviewSuggestionTextColor(prompt.DarkGreen),

		prompt.OptionSuggestionBGColor(prompt.DefaultColor),
		prompt.OptionDescriptionBGColor(prompt.DefaultColor),
		prompt.OptionSuggestionTextColor(prompt.DefaultColor),
		prompt.OptionDescriptionTextColor(prompt.DefaultColor),

		prompt.OptionSelectedSuggestionBGColor(prompt.DarkGreen),
		prompt.OptionSelectedDescriptionBGColor(prompt.DefaultColor),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGreen),
	)
	p.Run()

	j := jenny.Jenkins{
		Uri:      "uri",
		User:     "user",
		Password: "sd",
	}
	fmt.Println(j.PasswordMatch("sd"))
}
