package main

import "errors"

func myCode(input string) error {
        if len(input) == 0 { return errors.New("zero length") }
	if input[0] == '*' { return errors.New("no * expected") }
	if input[:3] == "foo" { return errors.New("foo is a reserved keyword") }
	return nil
}

func main() {
        myCode("")
        myCode("*")
        myCode("qck")
        myCode("qk")
}
