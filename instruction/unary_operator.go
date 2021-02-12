/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

type UnaryOperatorType uint8

const (
	UnaryOperatorBitwiseNot UnaryOperatorType = iota
	UnaryOperatorNot
)

func UnaryOperatorToOpCode(operator UnaryOperatorType) Commands {
	switch operator {
	case UnaryOperatorBitwiseNot:
		return CmdIntBitwiseNot
	case UnaryOperatorNot:
		return CmdBoolLogicalNot
	}

	panic("swamp opcodes: illegal unary operator")
}
