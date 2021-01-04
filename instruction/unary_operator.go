/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

type UnaryOperatorType uint8

const (
	UnaryOperatorBitwiseNot UnaryOperatorType = 0x01
	UnaryOperatorNot        UnaryOperatorType = 0x02
)

func UnaryOperatorToOpCode(operator UnaryOperatorType) Commands {
	switch operator {
	case UnaryOperatorBitwiseNot:
		return CmdBitwiseNot
	case UnaryOperatorNot:
		return CmdLogicalNot
	}

	panic("swamp opcodes: illegal unary operator")
}
