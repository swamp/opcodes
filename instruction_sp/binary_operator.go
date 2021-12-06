/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

// Package instruction_sp holds all the instructions for the Swamp runtime.
package instruction_sp

// BinaryOperatorType defines the type of binary operator.
type BinaryOperatorType uint8

// The binary operator types.
const (
	BinaryOperatorArithmeticIntPlus BinaryOperatorType = iota
	BinaryOperatorArithmeticIntMinus
	BinaryOperatorArithmeticIntDivide
	BinaryOperatorArithmeticIntRemainder
	BinaryOperatorArithmeticIntMultiply
	BinaryOperatorBooleanIntEqual
	BinaryOperatorBooleanIntNotEqual
	BinaryOperatorBooleanIntLess
	BinaryOperatorBooleanIntLessOrEqual
	BinaryOperatorBooleanIntGreater
	BinaryOperatorBooleanIntGreaterOrEqual
	BinaryOperatorBitwiseIntAnd
	BinaryOperatorBitwiseIntOr
	BinaryOperatorBitwiseIntXor
	BinaryOperatorArithmeticListAppend
	BinaryOperatorArithmeticFixedDivide
	BinaryOperatorArithmeticFixedMultiply
	BinaryOperatorBooleanStringEqual
	BinaryOperatorBooleanStringNotEqual
	BinaryOperatorBooleanEnumEqual
	BinaryOperatorBooleanEnumNotEqual
	BinaryOperatorBooleanBooleanEqual
	BinaryOperatorBooleanBooleanNotEqual
	BinaryOperatorBitwiseShiftLeft
	BinaryOperatorBitwiseShiftRight
	BinaryOperatorBitwiseNot
)

// BinaryOperatorToOpCode converts from the type of binary operator to the actual opcode instruction.
func BinaryOperatorToOpCode(operator BinaryOperatorType) Commands {
	switch operator {
	case BinaryOperatorArithmeticIntPlus:
		return CmdIntAdd
	case BinaryOperatorArithmeticIntMinus:
		return CmdIntSub
	case BinaryOperatorArithmeticIntDivide:
		return CmdIntDiv
	case BinaryOperatorArithmeticIntRemainder:
		return CmdIntRemainder
	case BinaryOperatorArithmeticIntMultiply:
		return CmdIntMul
	case BinaryOperatorArithmeticFixedDivide:
		return CmdFixedDiv
	case BinaryOperatorArithmeticFixedMultiply:
		return CmdFixedMul
	case BinaryOperatorArithmeticListAppend:
		return CmdListAppend
	case BinaryOperatorBooleanIntEqual:
		return CmdIntEqual
	case BinaryOperatorBooleanIntNotEqual:
		return CmdIntNotEqual
	case BinaryOperatorBooleanEnumEqual:
		return CmdEnumEqual
	case BinaryOperatorBooleanEnumNotEqual:
		return CmdEnumNotEqual
	case BinaryOperatorBooleanBooleanEqual:
		return CmdBoolEqual
	case BinaryOperatorBooleanBooleanNotEqual:
		return CmdBoolNotEqual
	case BinaryOperatorBooleanIntLess:
		return CmdIntLess
	case BinaryOperatorBooleanIntLessOrEqual:
		return CmdIntLessOrEqual
	case BinaryOperatorBooleanIntGreater:
		return CmdIntGreater
	case BinaryOperatorBooleanIntGreaterOrEqual:
		return CmdIntGreaterOrEqual
	case BinaryOperatorBitwiseIntAnd:
		return CmdIntBitwiseAnd
	case BinaryOperatorBitwiseIntOr:
		return CmdIntBitwiseOr
	case BinaryOperatorBitwiseIntXor:
		return CmdIntBitwiseXor
	case BinaryOperatorBitwiseShiftLeft:
		return CmdIntBitwiseShiftLeft
	case BinaryOperatorBitwiseShiftRight:
		return CmdIntBitwiseShiftLeft
	case BinaryOperatorBitwiseNot:
		return CmdIntBitwiseNot
	}

	panic("swamp opcodes: unknown binary operator")
}

// BinaryStringOperatorToOpCode converts from the type of binary operator to the actual opcode instruction.
func BinaryStringOperatorToOpCode(operator BinaryOperatorType) Commands {
	switch operator {
	case BinaryOperatorBooleanStringEqual:
		return CmdStringEqual
	case BinaryOperatorBooleanStringNotEqual:
		return CmdStringNotEqual
	}

	panic("swamp opcodes: unknown string binary operator")
}

// BinaryEnumOperatorToOpCode converts from the type of binary operator to the actual opcode instruction.
func BinaryEnumOperatorToOpCode(operator BinaryOperatorType) Commands {
	switch operator {
	case BinaryOperatorBooleanEnumEqual:
		return CmdEnumEqual
	case BinaryOperatorBooleanEnumNotEqual:
		return CmdEnumNotEqual
	}

	panic("swamp opcodes: unknown enum binary operator")
}

// BinaryEnumOperatorToOpCode converts from the type of binary operator to the actual opcode instruction.
func BinaryBooleanOperatorToOpCode(operator BinaryOperatorType) Commands {
	switch operator {
	case BinaryOperatorBooleanBooleanEqual:
		return CmdBoolEqual
	case BinaryOperatorBooleanBooleanNotEqual:
		return CmdBoolNotEqual
	}

	panic("swamp opcodes: unknown boolean binary operator")
}
