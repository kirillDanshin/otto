package parser

import "bytes"

type _RegExp_parser struct {
	str    string
	length int

	chr       rune // The current character
	chrOffset int  // The offset of current character
	offset    int  // The offset after current character (may be greater than 1)

	errors  []error
	invalid bool // The input is an invalid JavaScript RegExp

	goRegexp *bytes.Buffer
}
