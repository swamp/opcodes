/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type StringAppend struct {
	a           swampopcodetype.Register
	b           swampopcodetype.Register
	destination swampopcodetype.Register
}

func (c *StringAppend) Write(writer OpcodeWriter) error {
	writer.Command(CmdStringAppend)
	writer.Register(c.destination)
	writer.Register(c.a)
	writer.Register(c.b)

	return nil
}

func NewStringAppend(destination swampopcodetype.Register, a swampopcodetype.Register, b swampopcodetype.Register) *StringAppend {
	return &StringAppend{destination: destination, a: a, b: b}
}

func (c *StringAppend) String() string {
	return fmt.Sprintf("stringappend %v,%v,%v", c.destination, c.a, c.b)
}
