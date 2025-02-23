package parser

import "math"

func parseNumber(lit string) any {
	decimal := int64(0)
	fractional := int64(0)
	charCounter := 0
	exponent := int(0)

	decimalEnd := len(lit) - 1
	for i, s := range lit {
		if isDigit(s) {
			if i != 0 {
				decimal = (decimal * 10) // move decimal place
			}
			decimal += int64(s - '0') // add digit
		} else {
			decimalEnd = i
			charCounter++
			break
		}
	}
	if decimalEnd == len(lit)-1 {
		return int64(decimal)
	}
	fractionalEnd := decimalEnd
	if lit[decimalEnd] == '.' {
		if decimalEnd < len(lit)-1 { //we expect at least one character after . - .0
			// parse fractional
			for i, s := range lit[decimalEnd+charCounter:] {
				if isDigit(s) {
					if i != 0 {
						fractional = (fractional * 10)
					}
					fractional = (fractional + int64(s-'0'))
					fractionalEnd = i + 1 + decimalEnd
				} else {
					charCounter++
					fractionalEnd = i + decimalEnd
					break
				}
			}
		} else {
			fractionalEnd = decimalEnd // we have only . with no numbers after
		}
	}
	if fractionalEnd != len(lit)-1 && len(lit) > fractionalEnd+charCounter &&
		(lit[fractionalEnd+charCounter-1] == 'e' || lit[fractionalEnd+charCounter-1] == 'E') {
		if fractionalEnd < len(lit)-1 {
			for i, s := range lit[fractionalEnd+charCounter:] {
				if isDigit(s) {
					if i != 0 {
						exponent = (exponent * 10)
					}
					exponent = exponent + int(s-'0')
				}
			}
		}
	}
	if fractional == 0 && exponent == 0 {
		return decimal
	} else {
		f := float64(decimal)
		moveBy := float64(int(math.Pow(10, (float64(fractionalEnd) - float64(decimalEnd)))))
		decimal := float64(fractional) / moveBy
		f = f + decimal

		if exponent != 0 {
			f = math.Pow(f, float64(exponent))
		}
		return f
	}
}
