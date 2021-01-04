/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type ListConj struct {
	list        swampopcodetype.Register
	destination swampopcodetype.Register
	item        swampopcodetype.Register
}

func (c *ListConj) Write(writer OpcodeWriter) error {
	writer.Command(CmdListConj)
	writer.Register(c.destination)
	writer.Register(c.list)
	writer.Register(c.item)

	return nil
}

func NewListConj(destination swampopcodetype.Register, item swampopcodetype.Register, list swampopcodetype.Register) *ListConj {
	return &ListConj{destination: destination, item: item, list: list}
}

func (c *ListConj) String() string {
	return fmt.Sprintf("conj %v %v %v", c.destination, c.item, c.list)
}
