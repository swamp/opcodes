/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package instruction_sp

import (
	"fmt"

	opcode_sp_type "github.com/swamp/opcodes/type"
)

type MemoryCopy struct {
	a           opcode_sp_type.SourceStackPositionRange
	destination opcode_sp_type.TargetStackPosition
}

func (c *MemoryCopy) Write(writer OpcodeWriter) error {
	writer.Command(CmdCopyMemory)
	writer.TargetStackPosition(c.destination)
	writer.SourceStackPositionRange(c.a)

	return nil
}

func NewMemoryCopy(destination opcode_sp_type.TargetStackPosition,
	a opcode_sp_type.SourceStackPositionRange) *MemoryCopy {
	if a.Range == 0 {
		panic("zero range not allowed")
	}
	return &MemoryCopy{destination: destination, a: a}
}

func (c *MemoryCopy) String() string {
	return fmt.Sprintf("%s %v,%v", OpcodeToMnemonic(CmdCopyMemory), c.destination, c.a)
}
