/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package instruction_sp

import (
	"fmt"

	opcode_sp_type "github.com/swamp/opcodes/type"
)

type ShortRune uint8

func (r ShortRune) String() string {
	return fmt.Sprintf("'%c' (%d)", uint8(r), uint8(r))
}

type LoadRune struct {
	destination opcode_sp_type.TargetStackPosition
	shortRune   ShortRune
}

func (c *LoadRune) Write(writer OpcodeWriter) error {
	writer.Command(CmdLoadRune)
	writer.TargetStackPosition(c.destination)
	writer.Rune(c.shortRune)

	return nil
}

func NewLoadRune(destination opcode_sp_type.TargetStackPosition,
	shortRune ShortRune) *LoadRune {
	return &LoadRune{destination: destination, shortRune: shortRune}
}

func (c *LoadRune) String() string {
	return fmt.Sprintf("%s %v,%v", OpcodeToMnemonic(CmdLoadRune), c.destination, c.shortRune)
}
