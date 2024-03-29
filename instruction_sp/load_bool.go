/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package instruction_sp

import (
	"fmt"

	opcode_sp_type "github.com/swamp/opcodes/type"
)

type LoadBool struct {
	destination opcode_sp_type.TargetStackPosition
	a           bool
}

func (c *LoadBool) Write(writer OpcodeWriter) error {
	writer.Command(CmdLoadBoolean)
	writer.TargetStackPosition(c.destination)
	writer.Boolean(c.a)

	return nil
}

func NewLoadBool(destination opcode_sp_type.TargetStackPosition,
	a bool) *LoadBool {
	return &LoadBool{destination: destination, a: a}
}

func (c *LoadBool) String() string {
	return fmt.Sprintf("%s %v,%v", OpcodeToMnemonic(CmdLoadBoolean), c.destination, c.a)
}
