/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	swampopcodetype "github.com/swamp/opcodes/type"
)

// BoolNot inverts the specified register and puts the result in the destination register.
type BoolNot struct {
	source swampopcodetype.Register
	target swampopcodetype.Register
}

func NewBoolNot(target swampopcodetype.Register, source swampopcodetype.Register) *BoolNot {
	return &BoolNot{target: target, source: source}
}

func (c *BoolNot) Write(writer OpcodeWriter) error {
	writer.Command(CmdBoolLogicalNot)
	writer.Register(c.target)
	writer.Register(c.source)
	return nil
}

func (c *BoolNot) String() string {
	return "boolnot"
}
