package medium

const (
	MaxInt32 = 1<<31 - 1 // 2147483647
	MinInt32 = -1 << 31  // -2147483648
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func myAtoi(s string) int {
	var found int64
	sign := 1
	firstCharFound := false
	firstDigitFound := false
	for i, c := range s {
		if c == ' ' && !firstCharFound {
			continue
		} else {
			firstCharFound = true
		}
		if isDigit(c) {
			if !firstDigitFound && i > 0 && s[i-1] == '-' {
				sign = -1
			}
			firstDigitFound = true

			if found != 0 {
				found *= 10
			}
			found += int64(c - '0')

			if sign < 0 && found*int64(sign) < MinInt32 {
				return MinInt32
			} else if sign > 0 && found > MaxInt32 {
				return MaxInt32
			}
		} else if (c == '-' || c == '+') && !firstDigitFound {
			if i+1 >= len(s) || !isDigit(rune(s[i+1])) {
				break
			}
			continue
		} else {
			break
		}

	}
	return int(found) * sign
}
