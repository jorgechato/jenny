package jenny

type Jenkins struct {
	Uri      string
	User     string
	Password string
}

func (j *Jenkins) PasswordMatch(password string) bool {
	hash := GetMD5Hash(password)

	if hash == j.Password {
		return true
	}

	return false
}
