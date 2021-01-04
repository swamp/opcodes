/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type GetStruct struct {
	source      swampopcodetype.Register
	target      swampopcodetype.Register
	fieldLookup []swampopcodetype.Field
}

func (c *GetStruct) Write(writer OpcodeWriter) error {
	writer.Command(CmdStructGet)
	writer.Register(c.target)
	writer.Register(c.source)
	writer.Count(len(c.fieldLookup))

	for _, argument := range c.fieldLookup {
		writer.Field(argument)
	}

	return nil
}

func NewGetStruct(target swampopcodetype.Register, source swampopcodetype.Register, fieldLookup []swampopcodetype.Field) *GetStruct {
	return &GetStruct{target: target, source: source, fieldLookup: fieldLookup}
}

func (c *GetStruct) String() string {
	return fmt.Sprintf("get %v, %v, %v", c.target, c.source, c.fieldLookup)
}
