package main

func myAtoi(s string) int {
	const IntMax = 1<<31 - 1
	const IntMin = -1 << 31

	i, n := 0, len(s)

	for i < n && s[i] == ' ' {
		i++
	}

	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		if result > IntMax/10 || (result == IntMax/10 && digit > IntMax%10) {
			if sign == 1 {
				return IntMax
			}
			return IntMin
		}
		result = result*10 + digit
		i++
	}

	return sign * result
}