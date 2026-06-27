package easy

type RomanValue struct {
	Value           int
	BeforeLow       string
	BeforeLowValue  int
	BeforeHigh      string
	BeforeHighValue int
}

func romanToInt(s string) int {
	lib := map[string]RomanValue{
		"I": {
			Value:           1,
			BeforeLow:       "V",
			BeforeLowValue:  4,
			BeforeHigh:      "X",
			BeforeHighValue: 9,
		},
		"V": {
			Value: 5,
		},
		"X": {
			Value:           10,
			BeforeLow:       "L",
			BeforeLowValue:  40,
			BeforeHigh:      "C",
			BeforeHighValue: 90,
		},
		"L": {
			Value: 50,
		},
		"C": {
			Value:           100,
			BeforeLow:       "D",
			BeforeLowValue:  400,
			BeforeHigh:      "M",
			BeforeHighValue: 900,
		},
		"D": {
			Value: 500,
		},
		"M": {
			Value: 1000,
		},
	}

	res := 0
	l := len(s)
	for i := 0; i < l; i++ {
		cur := s[i : i+1]
		rv := lib[cur]
		val := rv.Value
		if i+1 < l {
			next := s[i+1 : i+2]
			if next == rv.BeforeLow {
				val = rv.BeforeLowValue
				i++
			}
			if next == rv.BeforeHigh {
				val = rv.BeforeHighValue
				i++
			}
		}
		res += val
	}

	return res
}
