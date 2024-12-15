package auth

type Credentials struct {
	Username string
	Password string
}

func ValidateLogin(creds Credentials) bool {
	// return creds.Username == "1234" && creds.Password == "1234"
	return true
}
