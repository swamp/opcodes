/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodetype

import "fmt"

type DeltaPC uint8

type ProgramCounter struct {
	position uint16
}

func (p ProgramCounter) String() string {
	return fmt.Sprintf("@%02x", p.position)
}

func NewProgramCounter(position uint16) ProgramCounter {
	return ProgramCounter{position: position}
}

func (p ProgramCounter) Add(delta uint16) ProgramCounter {
	return ProgramCounter{position: p.position + delta}
}

func (p ProgramCounter) Value() uint16 {
	return p.position
}

func (p ProgramCounter) IsAfter(other ProgramCounter) bool {
	return p.position > other.position
}

func (p ProgramCounter) Delta(after ProgramCounter) DeltaPC {
	return DeltaPC(after.position - p.position)
}
