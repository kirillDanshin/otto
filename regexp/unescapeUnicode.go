package regexp

import (
	"unicode/utf8"

	"github.com/kirillDanshin/myutils"
)

const (
	sequencePrefix = "x{"
	sequenceEnd    = "}"
)

func UnescapeUnicode(pattern string) string {
	s := []byte(pattern)
	result := ""
	for utf8.RuneCount(s) >= 1 {
		r, size := utf8.DecodeRune(s)
		s = s[size:]
		nextR, size := utf8.DecodeRune(s)
		switch r {
		case '\\':
			if nextR == 'u' || nextR == 'U' {
				result = myutils.Concat(result, uppercaseHex(r), sequencePrefix)
				s = s[size:]
				r, size = utf8.DecodeRune(s)
				result = myutils.Concat(result, uppercaseHex(r))

				s = s[size:]
				r, size = utf8.DecodeRune(s)
				result = myutils.Concat(result, uppercaseHex(r))

				s = s[size:]
				r, size = utf8.DecodeRune(s)
				result = myutils.Concat(result, uppercaseHex(r))

				s = s[size:]
				r, size = utf8.DecodeRune(s)
				result = myutils.Concat(result, uppercaseHex(r), sequenceEnd)

				s = s[size:]

			}

		default:
			result = myutils.Concat(result, string(r))
		}
	}
	return result
}

const (
	upperA = "A"
	upperB = "B"
	upperC = "C"
	upperD = "D"
	upperE = "E"
	upperF = "F"
)

func uppercaseHex(r rune) string {
	switch r {
	case 'a', 'A':
		return upperA
	case 'b', 'B':
		return upperB
	case 'c', 'C':
		return upperC
	case 'd', 'D':
		return upperD
	case 'e', 'E':
		return upperE
	case 'f', 'F':
		return upperF
	}
	return string(r)
}
