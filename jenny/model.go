package jenny

//Jenkins struct
type Jenkins struct {
	Project  string `yaml:"project"`
	Uri      string `yaml:"uri"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var jenkins = Jenkins{}

func (j *Jenkins) isEmpty() bool {
	if j.User == "" || j.Uri == "" || j.Password == "" {
		return true
	}
	return false
}

func (j *Jenkins) isProject() bool {
	return j.Project != ""
}
