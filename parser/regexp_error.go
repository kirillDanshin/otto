package parser

import "fmt"

// TODO Better error reporting, use the offset, etc.
func (self *_RegExp_parser) error(offset int, msg string, msgValues ...interface{}) error {
	err := fmt.Errorf(msg, msgValues...)
	self.errors = append(self.errors, err)
	return err
}
