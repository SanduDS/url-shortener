package base62

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(num int) string {
	if num == 0 {
		return string(charset[0])
	}

	result := ""
	base := len(charset)
	for num > 0 {
		remainder := num % base
		result = string(charset[remainder]) + result
		num = num / base
	}

	return result
}
