package jenny

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var filename string

func Init(msg bool) {
	if !isConfigured() && msg {
		color.Red("No %s found please type profile.\n", filename)
	} else {
		jenkins.login()
	}
}

func isConfigured() bool {
	jtmp, err := readYaml(jenkins)
	jenkins = jtmp
	return !err
}

func writeYaml(j Jenkins, force bool) {
	//struct to yaml
	if !force {
		j.Password = ""
		j.User = ""
	}
	d, err := yaml.Marshal(&j)
	check(err)
	//write file
	err = ioutil.WriteFile(filename, d, 0644)
	check(err)

	return
}

func removeYaml() {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		err := os.Remove(filename)
		check(err)
	}
}

func readYaml(j Jenkins) (Jenkins, bool) {
	cerr := isFile()

	if !cerr {
		yamlFile, err := ioutil.ReadFile(filename)
		check(err)

		err = yaml.Unmarshal(yamlFile, &j)
		check(err)
	}
	return j, cerr
}

func isFile() bool {
	dir, e := os.Getwd()
	check(e)
	currentFile := fmt.Sprintf("%s/.jenny.yml", dir)
	filename = fmt.Sprintf("%s/.jenny.yml", os.Getenv("HOME"))

	if _, err := os.Stat(currentFile); !os.IsNotExist(err) {
		filename = currentFile
		return false
	} else if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return false
	}

	return true
}

func setFilename(global bool) {
	dir, e := os.Getwd()
	check(e)
	currentFile := fmt.Sprintf("%s/.jenny.yml", dir)

	if global {
		filename = fmt.Sprintf("%s/.jenny.yml", os.Getenv("HOME"))
	} else {
		filename = currentFile
	}
}

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func printJenkins(j Jenkins, u bool) {
	tmp := j
	if !u {
		tmp.Password = "******"
	}

	d, err := yaml.Marshal(&tmp)
	check(err)
	fmt.Printf("--- Profile:\n%s", string(d))
}
