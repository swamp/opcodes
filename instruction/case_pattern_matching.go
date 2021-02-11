/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type CasePatternMatchingJump struct {
	literal swampopcodetype.Register
	jump    *swampopcodetype.Label
}

func NewCasePatternMatchingJump(literal swampopcodetype.Register, jump *swampopcodetype.Label) CasePatternMatchingJump {
	return CasePatternMatchingJump{literal: literal, jump: jump}
}

func (e CasePatternMatchingJump) String() string {
	return fmt.Sprintf("[%v %v]", e.literal, e.jump)
}

type CasePatternMatching struct {
	destination swampopcodetype.Register
	source      swampopcodetype.Register
	jumps       []CasePatternMatchingJump
}

func NewCasePatternMatching(destination swampopcodetype.Register, source swampopcodetype.Register,
	jumps []CasePatternMatchingJump) *CasePatternMatching {
	return &CasePatternMatching{destination: destination, source: source, jumps: jumps}
}

func (c *CasePatternMatching) Write(writer OpcodeWriter) error {
	writer.Command(CmdCasePatternMatching)
	writer.Register(c.destination)
	writer.Register(c.source)

	writer.Count(len(c.jumps))

	var lastLabel *swampopcodetype.Label

	for _, jump := range c.jumps {
		writer.Register(jump.literal)

		if lastLabel != nil {
			writer.LabelWithOffset(jump.jump, lastLabel)
		} else {
			writer.Label(jump.jump)
		}

		lastLabel = jump.jump
	}

	return nil
}

func (c *CasePatternMatching) String() string {
	return fmt.Sprintf("csep %v %v %v", c.destination, c.source, c.jumps)
}
