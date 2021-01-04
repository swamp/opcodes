/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type IntBinaryOperator struct {
	opcode      Commands
	a           swampopcodetype.Register
	b           swampopcodetype.Register
	destination swampopcodetype.Register
}

func (c *IntBinaryOperator) Write(writer OpcodeWriter) error {
	writer.Command(c.opcode)
	writer.Register(c.destination)
	writer.Register(c.a)
	writer.Register(c.b)

	return nil
}

func NewIntBinaryOperator(opcode Commands, destination swampopcodetype.Register, a swampopcodetype.Register, b swampopcodetype.Register) *IntBinaryOperator {
	return &IntBinaryOperator{opcode: opcode, destination: destination, a: a, b: b}
}

func (c *IntBinaryOperator) String() string {
	return fmt.Sprintf("%s %v,%v,%v", OpcodeToName(c.opcode), c.destination, c.a, c.b)
}
