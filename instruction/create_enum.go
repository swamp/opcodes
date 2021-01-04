/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type Enum struct {
	target         swampopcodetype.Register
	enumFieldIndex int
	arguments      []swampopcodetype.Register
}

func NewEnum(target swampopcodetype.Register, enumFieldIndex int, arguments []swampopcodetype.Register) *Enum {
	return &Enum{target: target, enumFieldIndex: enumFieldIndex, arguments: arguments}
}

func (c *Enum) Write(writer OpcodeWriter) error {
	writer.Command(CmdCreateEnum)
	writer.Register(c.target)
	writer.Count(c.enumFieldIndex)
	writer.Count(len(c.arguments))

	for _, argument := range c.arguments {
		writer.Register(argument)
	}

	return nil
}

func (c *Enum) String() string {
	return fmt.Sprintf("createenum %v %v (%v)", c.target, c.enumFieldIndex, c.arguments)
}
