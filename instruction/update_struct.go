/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type CopyToFieldInfo struct {
	Target swampopcodetype.Field
	Source swampopcodetype.Register
}

func (c CopyToFieldInfo) String() string {
	return fmt.Sprintf("%v<-%v", c.Target, c.Source)
}

type UpdateStruct struct {
	target           swampopcodetype.Register
	source           swampopcodetype.Register
	fieldDefinitions []CopyToFieldInfo
}

func NewUpdateStruct(target swampopcodetype.Register, source swampopcodetype.Register, fieldDefinitions []CopyToFieldInfo) *UpdateStruct {
	return &UpdateStruct{target: target, source: source, fieldDefinitions: fieldDefinitions}
}

func (c *UpdateStruct) Write(writer OpcodeWriter) error {
	writer.Command(CmdUpdateStruct)
	writer.Register(c.target)
	writer.Register(c.source)

	writer.Count(len(c.fieldDefinitions))

	for _, fieldDefinition := range c.fieldDefinitions {
		writer.Field(fieldDefinition.Target)
		writer.Register(fieldDefinition.Source)
	}

	return nil
}

func (c *UpdateStruct) String() string {
	return fmt.Sprintf("update %v %v %v", c.target, c.source, c.fieldDefinitions)
}
