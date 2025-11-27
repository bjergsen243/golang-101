package l18

func isValidPassword(password string) bool {

	if len(password) < 5 || len(password) > 12 {
		return false
	}
	hasLetter := false
	hasDigit := false
	for _, ch := range password {
		if ch >= 'A' && ch <= 'Z' {
			hasLetter = true
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}

	}
	return hasDigit && hasLetter
}
