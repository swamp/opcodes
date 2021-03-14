/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type StructSplit struct {
	source       swampopcodetype.Register
	destinations []swampopcodetype.Register
}

func NewStructSplit(source swampopcodetype.Register, destinations []swampopcodetype.Register) *StructSplit {
	return &StructSplit{source: source, destinations: destinations}
}

func (c *StructSplit) Write(writer OpcodeWriter) error {
	writer.Command(CmdStructSplit)
	writer.Register(c.source)
	writer.Count(len(c.destinations))

	for _, argument := range c.destinations {
		writer.Register(argument)
	}

	return nil
}

func (c *StructSplit) String() string {
	return fmt.Sprintf("stsp %v %v", c.source, c.destinations)
}
