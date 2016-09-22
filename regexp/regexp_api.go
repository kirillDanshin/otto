package regexp

import "github.com/kirillDanshin/go-pcre"

func (re *Regexp) FindStringIndex(s string) []int {
	return re.re.FindIndex([]byte(s), pcre.CONFIG_UTF32)
}

func (re *Regexp) FindStringSubmatchIndex(s string) (indexPairs []int) {
	r := re.FindAllStringIndex(s)
	for _, v := range r {
		indexPairs = append(indexPairs, v[0], v[1])
	}
	return
}

func (re *Regexp) FindAllStringSubmatchIndex(s string, _ int) [][]int {
	return re.re.FindAllIndex([]byte(s), pcre.CONFIG_UTF32)
}

func (re *Regexp) FindAllSubmatchIndex(b []byte, _ int) [][]int {
	result := make([][]int, len(b))

	for i := 0; i < len(b); {
		r := re.re.FindIndex(b[i:], 0)
		if r != nil {
			result = append(result, r)
			if len(r) >= 2 && r[1] != 0 {
				i += r[1]
				continue
			}
		}
		i++
	}
	return result
}

func (re *Regexp) FindAllStringIndex(str string, _ ...int) [][]int {
	return re.re.FindAllIndex([]byte(str), pcre.CONFIG_UTF32)
}

// IsCaseless checks if regexp is case insensetive
func (re *Regexp) IsCaseless() bool {
	return re.re.IsCaseless()
}

// IsMultiline checks if regexp is multiline
func (re *Regexp) IsMultiline() bool {
	return re.re.IsMultiline()
}

// FindAllIndex returns the start and end of the first match.
func (re *Regexp) FindAllIndex(bytes []byte, flags int) [][]int {
	return re.re.FindAllIndex(bytes, flags)
}

// FindIndex returns the start and end of the first match, or nil if no match.
// loc[0] is the start and loc[1] is the end.
func (re *Regexp) FindIndex(bytes []byte, flags int) []int {
	return re.re.FindIndex(bytes, flags)
}

// FindString returns the start and end of the first match, or nil if no match.
// loc[0] is the start and loc[1] is the end.
func (re *Regexp) FindString(s string, flags int) string {
	return re.re.FindString(s, flags)
}

// Groups returns the number of capture groups in the compiled regexp pattern.
func (re Regexp) Groups() int {
	return re.re.Groups()
}

// Match tries to match the speficied byte array slice to the current pattern.
// Returns true if the match succeeds.
func (re *Regexp) Match(subject []byte, flags int) bool {
	return re.re.Match(subject, flags)
}

// MatchString is same as Match, but accept string as argument
func (re *Regexp) MatchString(subject string, flags int) bool {
	return re.re.MatchString(subject, flags)
}

// Matcher returns a new matcher object, with the byte array slice as a
// subject.
func (re *Regexp) Matcher(subject []byte, flags int) (m *Matcher) {
	return &Matcher{re.re.Matcher(subject, flags)}
}

// MatcherString returns a new matcher object, with the specified subject string.
func (re *Regexp) MatcherString(subject string, flags int) (m *Matcher) {
	return &Matcher{re.re.MatcherString(subject, flags)}
}

// ReplaceAll returns a copy of a byte slice with pattern matches replaced by repl.
func (re *Regexp) ReplaceAll(bytes, repl []byte, flags int) []byte {
	return re.re.ReplaceAll(bytes, repl, flags)
}

// ReplaceAllString is same as ReplaceAll, but accept strings as arguments
func (re *Regexp) ReplaceAllString(src, repl string, flags int) string {
	return re.re.ReplaceAllString(src, repl, flags)
}
