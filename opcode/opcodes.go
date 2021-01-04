/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcode

import (
	"fmt"

	swampopcodeinst "github.com/swamp/opcodes/instruction"
	swampopcodetype "github.com/swamp/opcodes/type"
)

type OpCode uint8

type Count uint8

func DeltaProgramCounter(after swampopcodetype.ProgramCounter, before swampopcodetype.ProgramCounter) swampopcodetype.DeltaPC {
	if before.IsAfter(after) {
		panic(fmt.Sprintf("swamp opcodes: illegal delta program counter %v %v", before, after))
	}

	delta := before.Delta(after)
	return delta
}

type Instruction interface {
	Write(writer swampopcodeinst.OpcodeWriter) error
}
