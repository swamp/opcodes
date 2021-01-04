/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type CreateList struct {
	destination swampopcodetype.Register
	arguments   []swampopcodetype.Register
}

func NewCreateList(destination swampopcodetype.Register, arguments []swampopcodetype.Register) *CreateList {
	return &CreateList{destination: destination, arguments: arguments}
}

func (c *CreateList) Write(writer OpcodeWriter) error {
	writer.Command(CmdCreateList)
	writer.Register(c.destination)
	writer.Count(len(c.arguments))

	for _, argument := range c.arguments {
		writer.Register(argument)
	}

	return nil
}

func (c *CreateList) String() string {
	return fmt.Sprintf("crl %v %v", c.destination, c.arguments)
}
