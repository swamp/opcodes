/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type BinaryOperator struct {
	opcode      Commands
	a           swampopcodetype.Register
	b           swampopcodetype.Register
	destination swampopcodetype.Register
}

func (c *BinaryOperator) Write(writer OpcodeWriter) error {
	writer.Command(c.opcode)
	writer.Register(c.destination)
	writer.Register(c.a)
	writer.Register(c.b)

	return nil
}

func NewBinaryOperator(opcode Commands, destination swampopcodetype.Register, a swampopcodetype.Register, b swampopcodetype.Register) *BinaryOperator {
	return &BinaryOperator{opcode: opcode, destination: destination, a: a, b: b}
}

func (c *BinaryOperator) String() string {
	return fmt.Sprintf("%s %v,%v,%v", OpcodeToName(c.opcode), c.destination, c.a, c.b)
}
