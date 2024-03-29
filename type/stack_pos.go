/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package opcode_sp_type

import "fmt"

type StackPosition uint32

func (r StackPosition) String() string {
	return fmt.Sprintf("%v", uint32(r))
}

type (
	MemoryOffset uint32
	MemorySize   uint32
	MemoryAlign  uint32
)

const (
	Sizeof64BitPointer  MemorySize  = 8
	SizeofSwampInt      MemorySize  = 4
	SizeofSwampRune     MemorySize  = 4
	SizeofSwampBool     MemorySize  = 1

	AlignOfSwampBool = MemoryAlign(SizeofSwampBool)
	AlignOfSwampRune = MemoryAlign(SizeofSwampRune)
	AlignOfSwampInt  = MemoryAlign(SizeofSwampInt)
	Alignof64BitPointer MemoryAlign = 8
)

type StackRange uint16

type SourceStackRange StackRange

type TargetStackPosition StackPosition

func (r TargetStackPosition) String() string {
	return fmt.Sprintf("%v", uint32(r))
}

type SourceStackPosition StackPosition

func (r SourceStackPosition) String() string {
	return fmt.Sprintf("%v", uint32(r))
}

type ArgOffsetSize struct {
	Offset uint16
	Size   uint16
}

type ArgOffsetSizeAlign struct {
	Offset uint16
	Size   uint16
	Align  uint8
}

type StackPositionRange struct {
	Position StackPosition
	Range    StackRange
}

func (r StackPositionRange) String() string {
	return fmt.Sprintf("%v:%v", r.Position, r.Range)
}

type SourceStackPositionRange struct {
	Position SourceStackPosition
	Range    SourceStackRange
}

func (r SourceStackPositionRange) String() string {
	return fmt.Sprintf("(%v:%v)", r.Position, r.Range)
}

type TargetStackPositionRange StackPositionRange

func (r TargetStackPositionRange) String() string {
	return fmt.Sprintf("%v:%v<-", r.Position, r.Range)
}

type TargetFieldOffset StackRange

func (r TargetFieldOffset) String() string {
	return fmt.Sprintf("#%v", uint32(r))
}

type DynamicMemoryPosition uint32

type SourceDynamicMemoryPosition uint32

func (r SourceDynamicMemoryPosition) String() string {
	return fmt.Sprintf("$%04X", uint32(r))
}
