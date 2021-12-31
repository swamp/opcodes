/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package opcode_sp

import (
	"fmt"

	"github.com/swamp/opcodes/instruction_sp"
	opcode_sp_type "github.com/swamp/opcodes/type"
)

type Variable struct {
	pcStart *opcode_sp_type.Label
	pcEnd *opcode_sp_type.Label
	allocatedMemory opcode_sp_type.SourceStackPositionRange
}

func NewVariable(pcStart *opcode_sp_type.Label, allocatedMemory opcode_sp_type.SourceStackPositionRange) *Variable {
	return &Variable{pcStart: pcStart, allocatedMemory:allocatedMemory}
}

func (c *Variable) Write(writer instruction_sp.OpcodeWriter) error {
	return nil
}

func (c *Variable) DefineEnd(pcEnd *opcode_sp_type.Label) error {
	c.pcEnd = pcEnd
	return nil
}

func (c *Variable) String() string {
	return fmt.Sprintf("variable: %v", c.allocatedMemory)
}


type VariableEnd struct {
	refer *Variable
}

func NewVariableEnd(refer *Variable) *VariableEnd {
	return &VariableEnd{refer: refer}
}

func (c *VariableEnd) DefineEnd(endPosition *opcode_sp_type.Label) error {
	return c.refer.DefineEnd(endPosition)
}


func (c *VariableEnd) Write(writer instruction_sp.OpcodeWriter) error {
	return nil
}

func (c *VariableEnd) String() string {
	return fmt.Sprintf("variableend: %v", c.refer)
}
