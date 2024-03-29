/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package instruction_sp

import opcode_sp_type "github.com/swamp/opcodes/type"

type OpcodeWriter interface {
	SourceStackPosition(r opcode_sp_type.SourceStackPosition)
	TargetStackPosition(r opcode_sp_type.TargetStackPosition)
	SourceDynamicMemoryPosition(r opcode_sp_type.SourceDynamicMemoryPosition)
	Int32(r int32)
	Boolean(r bool)
	Rune(r ShortRune)
	// StackPositionRange(r opcode_sp_type.StackPositionRange)
	SourceStackPositionRange(r opcode_sp_type.SourceStackPositionRange)
	StackRange(r opcode_sp_type.StackRange)
	MemoryAlign(r opcode_sp_type.MemoryAlign)
	TargetFieldOffset(r opcode_sp_type.TargetFieldOffset)

	DeltaPC(pc opcode_sp_type.DeltaPC)
	Label(l *opcode_sp_type.Label)
	LabelWithOffset(l *opcode_sp_type.Label, offset *opcode_sp_type.Label)
	EnumValue(v uint8)
	Count(c int)
	ArgOffsetSize(opcode_sp_type.ArgOffsetSize)
	ArgOffsetSizeAlign(align opcode_sp_type.ArgOffsetSizeAlign)
	TypeIDConstant(c uint16)
	Command(cmd Commands)
}
