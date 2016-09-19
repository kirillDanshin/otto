package parser

func (self *_RegExp_parser) scan() {
	for self.chr != -1 {
		switch self.chr {
		case '\\':
			self.read()
			self.scanEscape(false)
		case '(':
			self.pass()
			self.scanGroup()
		case '[':
			self.pass()
			self.scanBracket()
		case ')':
			self.error(-1, "Unmatched ')'")
			self.invalid = true
			self.pass()
		default:
			self.pass()
		}
	}
}
