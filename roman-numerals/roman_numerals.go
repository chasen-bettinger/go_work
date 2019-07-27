package romannumerals

import "strconv"
import "errors"

var arabicToRoman = map[int]string{
	100: "C",
	200: "CC",
	300: "CCC",
	400: "CD",
	500: "D",
	600: "DC",
	700: "DCC",
	800: "DCCC",
	900: "CM",
	10:  "X",
	20:  "XX",
	30:  "XXX",
	40:  "XL",
	50:  "L",
	60:  "LX",
	70:  "LXX",
	80:  "LXXX",
	90:  "XC",
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
}

// ToRomanNumeral takes an arabic integer and transforms it into the roman numeral equivalent
func ToRomanNumeral(arabic int) (string, error) {
	romanNumerals := ""
	if arabic <= 0 || arabic > 3000 {
		return romanNumerals, errors.New("arabic integer cannot be less than or equal to 0 or greater than 3000")
	}
	arabicNumberAsString := strconv.Itoa(arabic)
	numberOfDigits := len(arabicNumberAsString)
	for _, val := range arabicNumberAsString {
		num := int(val - '0')
		if num == 0 {
			numberOfDigits--
			continue
		}
		if numberOfDigits == 4 {
			for num > 0 {
				romanNumerals += "M"
				num--
			}
		}
		if numberOfDigits == 3 {
			num *= 100
			romanNumerals += arabicToRoman[num]
		}
		if numberOfDigits == 2 {
			num *= 10
			romanNumerals += arabicToRoman[num]
		}
		if numberOfDigits == 1 {
			romanNumerals += arabicToRoman[num]
		}
		numberOfDigits--
	}

	return romanNumerals, nil
}
