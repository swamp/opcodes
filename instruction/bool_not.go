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
	a           swampopcodetype.Register
	destination swampopcodetype.Register
}

func NewBoolNot(destination swampopcodetype.Register, a swampopcodetype.Register) *BoolNot {
	return &BoolNot{destination: destination, a: a}
}

func (c *BoolNot) Write(writer OpcodeWriter) error {
	return nil
}

func (c *BoolNot) String() string {
	return "boolornot"
}
