package parser

import (
	"bytes"
	"fmt"
)

// TransformRegExp transforms a JavaScript pattern into  a Go "regexp" pattern.
//
// re2 (Go) cannot do backtracking, so the presence of a lookahead (?=) (?!) or
// backreference (\1, \2, ...) will cause an error.
//
// re2 (Go) has a different definition for \s: [\t\n\f\r ].
// The JavaScript definition, on the other hand, also includes \v, Unicode "Separator, Space", etc.
//
// If the pattern is invalid (not valid even in JavaScript), then this function
// returns the empty string and an error.
//
// If the pattern is valid, but incompatible (contains a lookahead or backreference),
// then this function returns the transformation (a non-empty string) AND an error.
func TransformRegExp(pattern string) (string, error) {

	if pattern == "" {
		return "", nil
	}

	fmt.Printf("Input: regexp=[%s]\n", pattern)

	// TODO If without \, if without (?=, (?!, then another shortcut

	parser := _RegExp_parser{
		str:      pattern,
		length:   len(pattern),
		goRegexp: bytes.NewBuffer(make([]byte, 0, 3*len(pattern)/2)),
	}
	parser.read() // Pull in the first character
	parser.scan()
	var err error
	if len(parser.errors) > 0 {
		err = parser.errors[0]
	}
	if parser.invalid {
		return "", err
	}

	fmt.Printf("Output: regexp=[%s] err=[%s]\n", parser.goRegexp.String(), err)
	// Might not be re2 compatible, but is still a valid JavaScript RegExp
	return parser.goRegexp.String(), err
}
