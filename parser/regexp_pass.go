package parser

func (self *_RegExp_parser) pass() {
	if self.chr != -1 {
		_, err := self.goRegexp.WriteRune(self.chr)
		if err != nil {
			self.errors = append(self.errors, err)
		}
	}
	self.read()
}
