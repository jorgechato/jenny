package jenny

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var filename = fmt.Sprintf("%s/.jenny.yml", os.Getenv("HOME"))

func Init() {
	if !IsConfigured() {
		fmt.Printf("No %s found please type profile.\n", filename)
	}
}

func IsConfigured() bool {
	jtmp = ReadYaml(jtmp)

	if jtmp.IsEmpty() {
		return false
	}
	jenkins = jtmp
	return true
}

func WriteYaml(j Jenkins) {
	if !j.IsEmpty() {
		//struct to yaml
		d, err := yaml.Marshal(&j)
		check(err)
		//write file
		//TODO: check if profile is stored and change it
		//TODO: keep the rest of profiles stored
		err = ioutil.WriteFile(filename, d, 0644)
		check(err)
	}
	return
}

//TODO: read specific profile, read only profile, read default profile
func ReadYaml(j Jenkins) Jenkins {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		yamlFile, err := ioutil.ReadFile(filename)
		check(err)
		//TODO: read a list of jenkins struct
		//m := make(map[interface{}]interface{})

		err = yaml.Unmarshal(yamlFile, &j)
		check(err)
	}
	return j
}

func FindProfile() {
}

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
		panic(e)
	}
}

func PrintJenkins(j Jenkins) {
	d, err := yaml.Marshal(&j)
	check(err)
	fmt.Printf("--- Profile:\n%s", string(d))
}
