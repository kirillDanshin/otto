package parser

// [...]
func (self *_RegExp_parser) scanBracket() {
	for self.chr != -1 {
		if self.chr == ']' {
			break
		} else if self.chr == '\\' {
			self.read()
			self.scanEscape(true)
			continue
		}
		self.pass()
	}
	if self.chr != ']' {
		self.error(-1, "Unterminated character class")
		self.invalid = true
		return
	}
	self.pass()
}
