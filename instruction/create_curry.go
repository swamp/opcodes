/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type Curry struct {
	target    swampopcodetype.Register
	function  swampopcodetype.Register
	arguments []swampopcodetype.Register
}

func NewCurry(target swampopcodetype.Register, function swampopcodetype.Register, arguments []swampopcodetype.Register) *Curry {
	return &Curry{target: target, function: function, arguments: arguments}
}

func (c *Curry) Write(writer OpcodeWriter) error {
	writer.Command(CmdCurry)
	writer.Register(c.target)
	writer.Register(c.function)
	writer.Count(len(c.arguments))

	for _, argument := range c.arguments {
		writer.Register(argument)
	}

	return nil
}

func (c *Curry) String() string {
	return fmt.Sprintf("curry %v %v (%v)", c.target, c.function, c.arguments)
}
