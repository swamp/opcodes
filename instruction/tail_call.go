/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import swampopcodetype "github.com/swamp/opcodes/type"

type TailCall struct {
	arguments []swampopcodetype.Register
}

func NewTailCall(arguments []swampopcodetype.Register) *TailCall {
	return &TailCall{arguments: arguments}
}

func (c *TailCall) Write(writer OpcodeWriter) error {
	writer.Command(CmdTailCall)

	for _, argument := range c.arguments {
		writer.Register(argument)
	}

	return nil
}
