package parser

import (
	"bytes"
	"log"
	"os"
	"runtime"

	"github.com/kirillDanshin/myutils"
)

const (
	_regexp_empty            = ""
	_regexp_any_nc           = ".*"
	_regexp_any_c            = "(.*)"
	_regexp_anyplus_nc       = ".+"
	_regexp_anyplus_c        = "(.+)"
	_regexp_squote_nc        = `'`
	_regexp_squote_c         = `(')`
	_regexp_dquote_nc        = `"`
	_regexp_dquote_c         = `(")`
	_regexp_plus_nc          = `\+`
	_regexp_plus_c           = `(\+)`
	_regexp_urlenc_space_nc  = `%20`
	_regexp_urlenc_space_c   = `(%20)`
	_regexp_http_validation  = `^https?:\/\/[^\/]*`
	_regexp_any_uppercase_nc = `[A-Z]`
	_regexp_any_uppercase_c  = `([A-Z])`
	_regexp_is_ms_pref       = `^ms-`
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

	log.Printf("runtime.Caller(1): %#+v\n", myutils.Slice(runtime.Caller(1))...)
	os.Exit(349857394)

	// if pattern == "" {
	// 	return "", nil
	// }

	switch pattern {
	// common cases that will leave as is
	// but frequently occurred in many
	// popular frameworks like react.js etc.
	case _regexp_empty,
		_regexp_any_nc,
		_regexp_any_c,
		_regexp_anyplus_nc,
		_regexp_anyplus_c,

		_regexp_squote_nc,
		_regexp_squote_c,
		_regexp_dquote_nc,
		_regexp_dquote_c,

		_regexp_plus_nc,
		_regexp_plus_c,

		_regexp_urlenc_space_nc,
		_regexp_urlenc_space_c,
		_regexp_http_validation,

		_regexp_any_uppercase_nc,
		_regexp_any_uppercase_c,

		_regexp_is_ms_pref:
		return pattern, nil
	}

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
		log.Printf("Input: regexp=[%s]\n", pattern)
		log.Printf("Output: regexp=[%s] err=[%s]\n", parser.goRegexp.String(), err)
		return "", err
	}

	// Might not be re2 compatible, but is still a valid JavaScript RegExp
	return parser.goRegexp.String(), err
}
