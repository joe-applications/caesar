package encrypt


func Encrypt(str string) string {
	// Encrypt string with right shift by 3
	var encryptedStr string
	for _, char := range str {
		encryptedStr += string(int(char) + 3)
	}
	return encryptedStr
}
