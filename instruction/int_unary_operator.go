/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type IntUnaryOperator struct {
	opcode      Commands
	a           swampopcodetype.Register
	destination swampopcodetype.Register
}

func NewIntUnaryOperator(opcode Commands, destination swampopcodetype.Register, a swampopcodetype.Register) *IntUnaryOperator {
	return &IntUnaryOperator{opcode: opcode, destination: destination, a: a}
}

func (c *IntUnaryOperator) String() string {
	return fmt.Sprintf("%s %v,%v", OpcodeToName(c.opcode), c.destination, c.a)
}

func (c *IntUnaryOperator) Write(writer OpcodeWriter) error {
	writer.Command(c.opcode)
	writer.Register(c.destination)
	writer.Register(c.a)

	return nil
}
