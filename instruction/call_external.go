/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

// CallExternal is an instruction that calls into the embedder (usually C/C++ code).
type CallExternal struct {
	target    swampopcodetype.Register
	function  swampopcodetype.Register
	arguments []swampopcodetype.Register
}

func NewCallExternal(target swampopcodetype.Register, function swampopcodetype.Register, arguments []swampopcodetype.Register) *CallExternal {
	return &CallExternal{target: target, function: function, arguments: arguments}
}

func (c *CallExternal) Write(writer OpcodeWriter) error {
	writer.Command(CmdCallExternal)
	writer.Register(c.target)
	writer.Register(c.function)
	writer.Count(len(c.arguments))

	for _, argument := range c.arguments {
		writer.Register(argument)
	}

	return nil
}

func (c *CallExternal) String() string {
	return fmt.Sprintf("callexternal %v %v (%v)", c.target, c.function, c.arguments)
}
