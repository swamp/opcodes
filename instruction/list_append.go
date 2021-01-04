/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type ListAppend struct {
	a           swampopcodetype.Register
	b           swampopcodetype.Register
	destination swampopcodetype.Register
}

func (c *ListAppend) Write(writer OpcodeWriter) error {
	writer.Command(CmdListAppend)
	writer.Register(c.destination)
	writer.Register(c.a)
	writer.Register(c.b)

	return nil
}

func NewListAppend(destination swampopcodetype.Register, a swampopcodetype.Register, b swampopcodetype.Register) *ListAppend {
	return &ListAppend{destination: destination, a: a, b: b}
}

func (c *ListAppend) String() string {
	return fmt.Sprintf("listappend %v,%v,%v", c.destination, c.a, c.b)
}
