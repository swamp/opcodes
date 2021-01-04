/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

type Commands uint8

const (
	CmdCreateStruct Commands = 0x01
	CmdUpdateStruct          = 0x02
	CmdStructGet             = 0x03
	CmdRegCopy               = 0x04
	CmdListConj              = 0x05

	CmdEnumCase    = 0x06
	CmdBranchFalse = 0x07
	CmdJump        = 0x08

	CmdCall         = 0x09
	CmdReturn       = 0x0a
	CmdCallExternal = 0x0b
	CmdTailCall     = 0x0c

	CmdAdd = 0x0d
	CmdSub = 0x0e
	CmdMul = 0x0f
	CmdDiv = 0x10

	CmdEqual          = 0x11
	CmdNotEqual       = 0x12
	CmdLess           = 0x13
	CmdLessOrEqual    = 0x14
	CmdGreater        = 0x15
	CmdGreaterOrEqual = 0x16
	CmdBitwiseAnd     = 0x17
	CmdBitwiseOr      = 0x18
	CmdBitwiseXor     = 0x19
	CmdBitwiseNot     = 0x1a
	CmdLogicalNot     = 0x1b
	CmdBranchTrue     = 0x1c

	CmdCurry        = 0x20
	CmdCreateList   = 0x21
	CmdListAppend   = 0x22
	CmdCreateEnum   = 0x23
	CmdStringAppend = 0x24

	CmdFixedMul = 0x25
	CmdFixedDiv = 0x26
)

func OpcodeToName(cmd Commands) string {
	names := map[Commands]string{
		CmdCreateStruct: "crs",
		CmdUpdateStruct: "upd",
		CmdStructGet:    "get",
		CmdEnumCase:     "case",
		CmdRegCopy:      "lr",
		CmdTailCall:     "tcl",
		CmdCall:         "call",
		CmdReturn:       "ret",
		CmdCallExternal: "ecall",
		CmdListConj:     "conj",
		CmdJump:         "jmp",
		CmdBranchFalse:  "bne",
		CmdAdd:          "add",
		CmdSub:          "sub",
		CmdMul:          "mul",
		CmdDiv:          "div",

		CmdEqual:          "cpeq",
		CmdNotEqual:       "cpne",
		CmdLess:           "cpl",
		CmdLessOrEqual:    "cple",
		CmdGreater:        "cpg",
		CmdGreaterOrEqual: "cpge",
		CmdBitwiseAnd:     "band",
		CmdBitwiseOr:      "bor",
		CmdBitwiseXor:     "bxor",
		CmdBitwiseNot:     "bnot",
		CmdLogicalNot:     "not",
		CmdCurry:          "curry",
		CmdCreateList:     "crl",
		CmdListAppend:     "lap",
		CmdFixedMul:       "fxmul",
		CmdFixedDiv:       "fxdiv",
	}

	return names[cmd]
}
