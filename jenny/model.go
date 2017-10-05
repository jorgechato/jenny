package jenny

type Jenkins struct {
	Uri      string
	User     string
	Password string
}

func (j *Jenkins) PasswordMatch(password string) bool {
	if password == j.Password {
		return true
	}

	return false
}
