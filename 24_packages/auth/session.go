package auth

func extractSession() string {
	return "LoggedIn"
}

func GetSession() string {
	return extractSession()
}
