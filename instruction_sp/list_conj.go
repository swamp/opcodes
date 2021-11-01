/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package instruction_sp

import (
	"fmt"

	opcode_sp_type "github.com/swamp/opcodes/type"
)

type ListConj struct {
	list        opcode_sp_type.SourceStackPosition
	destination opcode_sp_type.TargetStackPosition
	item        opcode_sp_type.SourceStackPosition
	itemSize    opcode_sp_type.StackRange
	itemAlign   opcode_sp_type.MemoryAlign
}

func (c *ListConj) Write(writer OpcodeWriter) error {
	writer.Command(CmdListConj)
	writer.TargetStackPosition(c.destination)
	writer.SourceStackPosition(c.list)
	writer.SourceStackPosition(c.item)
	writer.StackRange(c.itemSize)
	writer.MemoryAlign(c.itemAlign)

	return nil
}

func NewListConj(destination opcode_sp_type.TargetStackPosition, item opcode_sp_type.SourceStackPosition,
	itemSize opcode_sp_type.StackRange, itemAlign opcode_sp_type.MemoryAlign, list opcode_sp_type.SourceStackPosition) *ListConj {
	return &ListConj{destination: destination, item: item, list: list}
}

func (c *ListConj) String() string {
	return fmt.Sprintf("conj %v %v %v (%d, %d)", c.destination, c.item, c.list, c.itemSize, c.itemAlign)
}
