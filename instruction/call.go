/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type Call struct {
	target    swampopcodetype.Register
	function  swampopcodetype.Register
	arguments []swampopcodetype.Register
}

func NewCall(target swampopcodetype.Register, function swampopcodetype.Register, arguments []swampopcodetype.Register) *Call {
	return &Call{target: target, function: function, arguments: arguments}
}

func (c *Call) Write(writer OpcodeWriter) error {
	writer.Command(CmdCall)
	writer.Register(c.target)
	writer.Register(c.function)
	writer.Count(len(c.arguments))

	for _, argument := range c.arguments {
		writer.Register(argument)
	}

	return nil
}

func (c *Call) String() string {
	return fmt.Sprintf("call %v %v (%v)", c.target, c.function, c.arguments)
}
