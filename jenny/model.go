package jenny

//import "gopkg.in/yaml.v2"

type Jenkins struct {
	Name     string `yaml:"name"`
	Uri      string `yaml:"uri"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var jenkins = Jenkins{}
var jtmp = Jenkins{Name: "Default"}

func (j *Jenkins) IsEmpty() bool {
	if j.Name == "" || j.User == "" || j.Uri == "" || j.Password == "" {
		return true
	}
	return false
}
