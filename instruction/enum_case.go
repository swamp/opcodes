/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type EnumCaseJump struct {
	enumValue uint8
	arguments []swampopcodetype.Register
	jump      *swampopcodetype.Label
}

func NewEnumCaseJump(enumValue uint8, arguments []swampopcodetype.Register, jump *swampopcodetype.Label) EnumCaseJump {
	return EnumCaseJump{enumValue: enumValue, arguments: arguments, jump: jump}
}

func (e EnumCaseJump) String() string {
	return fmt.Sprintf("[%v %v %v]", e.enumValue, e.arguments, e.jump)
}

type EnumCase struct {
	destination swampopcodetype.Register
	source      swampopcodetype.Register
	jumps       []EnumCaseJump
}

func NewEnumCase(destination swampopcodetype.Register, source swampopcodetype.Register, jumps []EnumCaseJump) *EnumCase {
	return &EnumCase{destination: destination, source: source, jumps: jumps}
}

func (c *EnumCase) Write(writer OpcodeWriter) error {
	writer.Command(CmdEnumCase)
	writer.Register(c.destination)
	writer.Register(c.source)

	writer.Count(len(c.jumps))

	var lastLabel *swampopcodetype.Label

	for _, jump := range c.jumps {
		writer.EnumValue(jump.enumValue)
		writer.Count(len(jump.arguments))

		for _, arg := range jump.arguments {
			writer.Register(arg)
		}

		if lastLabel != nil {
			writer.LabelWithOffset(jump.jump, lastLabel)
		} else {
			writer.Label(jump.jump)
		}

		lastLabel = jump.jump
	}

	return nil
}

func (c *EnumCase) String() string {
	return fmt.Sprintf("cse %v %v %v", c.destination, c.source, c.jumps)
}
