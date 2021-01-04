/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type CreateStruct struct {
	destination swampopcodetype.Register
	arguments   []swampopcodetype.Register
}

func NewCreateStruct(destination swampopcodetype.Register, arguments []swampopcodetype.Register) *CreateStruct {
	return &CreateStruct{destination: destination, arguments: arguments}
}

func (c *CreateStruct) Write(writer OpcodeWriter) error {
	writer.Command(CmdCreateStruct)
	writer.Register(c.destination)
	writer.Count(len(c.arguments))

	for _, argument := range c.arguments {
		writer.Register(argument)
	}

	return nil
}

func (c *CreateStruct) String() string {
	return fmt.Sprintf("crs %v %v", c.destination, c.arguments)
}
