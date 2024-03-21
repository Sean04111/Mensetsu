package util

// atoi 函数转化string 到十进制int
func atoi(s string) int {
	integer := 0
	intengerSlice := []rune{}
	neg := false
	if s[0] == '-' {
		neg = true
	}
	for _, ch := range s {
		chNum := ch - '0'
		if 0 <= chNum && chNum <= 9 {
			intengerSlice = append(intengerSlice, chNum)
		}
	}
	x := 1
	for i := len(intengerSlice) - 1; i >= 0; i-- {
		integer = int(intengerSlice[i])*x + integer
		x *= 10
	}
	if neg {
		return -integer
	}
	return integer
}
