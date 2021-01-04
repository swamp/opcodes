/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

// BranchTrue branches if the test register is True.
type BranchTrue struct {
	test swampopcodetype.Register
	jump *swampopcodetype.Label
}

// NewBranchTrue creates a new BranchTrue struct.
func NewBranchTrue(test swampopcodetype.Register, jump *swampopcodetype.Label) *BranchTrue {
	return &BranchTrue{test: test, jump: jump}
}

func (c *BranchTrue) Write(writer OpcodeWriter) error {
	writer.Command(CmdBranchTrue)
	writer.Register(c.test)
	writer.Label(c.jump)

	return nil
}

func (c *BranchTrue) String() string {
	return fmt.Sprintf("brne %v %v", c.test, c.jump)
}
