package decrypt


func Decrypt(str string) string {
	var decryptedStr string
	for _, char := range str {
		decryptedStr += string(int(char) - 3)
	}
	return decryptedStr
}
