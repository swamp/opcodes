/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

// Package swampopcodeinst holds all the instructions for the Swamp runtime.
package swampopcodeinst

// BinaryOperatorType defines the type of binary operator.
type BinaryOperatorType uint8

// The binary operator types.
const (
	BinaryOperatorArithmeticPlus          BinaryOperatorType = 0x01
	BinaryOperatorArithmeticMinus         BinaryOperatorType = 0x02
	BinaryOperatorArithmeticDivide        BinaryOperatorType = 0x03
	BinaryOperatorArithmeticMultiply      BinaryOperatorType = 0x04
	BinaryOperatorBooleanEqual            BinaryOperatorType = 0x05
	BinaryOperatorBooleanNotEqual         BinaryOperatorType = 0x06
	BinaryOperatorBooleanLess             BinaryOperatorType = 0x07
	BinaryOperatorBooleanLessOrEqual      BinaryOperatorType = 0x08
	BinaryOperatorBooleanGreater          BinaryOperatorType = 0x09
	BinaryOperatorBooleanGreaterOrEqual   BinaryOperatorType = 0x0a
	BinaryOperatorBitwiseAnd              BinaryOperatorType = 0x0b
	BinaryOperatorBitwiseOr               BinaryOperatorType = 0x0c
	BinaryOperatorBitwiseXor              BinaryOperatorType = 0x0d
	BinaryOperatorArithmeticAppend        BinaryOperatorType = 0x0e
	BinaryOperatorArithmeticFixedDivide   BinaryOperatorType = 0x0f
	BinaryOperatorArithmeticFixedMultiply BinaryOperatorType = 0x10
)

// BinaryOperatorToOpCode converts from the type of binary operator to the actual opcode instruction.
func BinaryOperatorToOpCode(operator BinaryOperatorType) Commands {
	switch operator {
	case BinaryOperatorArithmeticPlus:
		return CmdAdd
	case BinaryOperatorArithmeticMinus:
		return CmdSub
	case BinaryOperatorArithmeticDivide:
		return CmdDiv
	case BinaryOperatorArithmeticMultiply:
		return CmdMul
	case BinaryOperatorArithmeticFixedDivide:
		return CmdFixedDiv
	case BinaryOperatorArithmeticFixedMultiply:
		return CmdFixedMul
	case BinaryOperatorArithmeticAppend:
		return CmdListAppend
	case BinaryOperatorBooleanEqual:
		return CmdEqual
	case BinaryOperatorBooleanNotEqual:
		return CmdNotEqual
	case BinaryOperatorBooleanLess:
		return CmdLess
	case BinaryOperatorBooleanLessOrEqual:
		return CmdLessOrEqual
	case BinaryOperatorBooleanGreater:
		return CmdGreater
	case BinaryOperatorBooleanGreaterOrEqual:
		return CmdGreaterOrEqual
	case BinaryOperatorBitwiseAnd:
		return CmdBitwiseAnd
	case BinaryOperatorBitwiseOr:
		return CmdBitwiseOr
	case BinaryOperatorBitwiseXor:
		return CmdBitwiseXor
	}

	panic("swamp opcodes: unknown binary operator")
}
