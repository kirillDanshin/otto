package parser

import "log"

// (...)
func (self *_RegExp_parser) scanGroup() {
	// str := self.str[self.chrOffset:]
	// if len(str) > 1 { // A possibility of (?= or (?!
	// 	if str[0] == '?' {
	// 		if str[1] == '=' || str[1] == '!' {
	// 			self.error(-1, "re2: Invalid (%s) <lookahead>", self.str[self.chrOffset:self.chrOffset+2])
	// 		}
	// 	}
	// }
	if len(self.str[self.chrOffset:]) > 1 && self.str[self.chrOffset] == '?' {
		switch self.str[self.chrOffset+1] {
		case '=':
			log.Printf("(?=) occurred")
			self.error(-1, "re2: Invalid (%s) <lookahead>", self.str[self.chrOffset:self.chrOffset+2])
		case '!':
			log.Printf("(?!) occurred")
			self.error(-1, "re2: Invalid (%s) <lookahead>", self.str[self.chrOffset:self.chrOffset+2])
		}
	}
	for self.chr != -1 && self.chr != ')' {
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
		default:
			self.pass()
			continue
		}
	}
	if self.chr != ')' {
		self.error(-1, "Unterminated group")
		self.invalid = true
		return
	}
	self.pass()
}
