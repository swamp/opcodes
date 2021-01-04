/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type RegCopy struct {
	source swampopcodetype.Register
	target swampopcodetype.Register
}

func NewRegCopy(target swampopcodetype.Register, source swampopcodetype.Register) *RegCopy {
	return &RegCopy{target: target, source: source}
}

func (c *RegCopy) Write(writer OpcodeWriter) error {
	writer.Command(CmdRegCopy)
	writer.Register(c.target)
	writer.Register(c.source)

	return nil
}

func (c *RegCopy) String() string {
	return fmt.Sprintf("lr %v,%v", c.target, c.source)
}
