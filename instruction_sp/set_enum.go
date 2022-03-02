/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package instruction_sp

import (
	"fmt"

	opcode_sp_type "github.com/swamp/opcodes/type"
)

type SetEnum struct {
	destination opcode_sp_type.TargetStackPosition
	enumIndex   uint8
	itemSize    opcode_sp_type.StackRange
}

func (c *SetEnum) Write(writer OpcodeWriter) error {
	writer.Command(CmdSetEnum)
	writer.TargetStackPosition(c.destination)
	writer.EnumValue(c.enumIndex)
	writer.StackRange(c.itemSize)

	return nil
}

func NewSetEnum(destination opcode_sp_type.TargetStackPosition,
	enumIndex uint8, itemSize opcode_sp_type.StackRange) *SetEnum {
	return &SetEnum{destination: destination, enumIndex: enumIndex, itemSize: itemSize}
}

func (c *SetEnum) String() string {
	return fmt.Sprintf("%s %v,%v (%v)", OpcodeToMnemonic(CmdSetEnum), c.destination, c.enumIndex, c.itemSize)
}
