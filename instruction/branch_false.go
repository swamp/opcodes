/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

// BranchFalse branches if the test register is False.
type BranchFalse struct {
	test swampopcodetype.Register
	jump *swampopcodetype.Label
}

// NewBranchFalse creates a new BranchFalse struct.
func NewBranchFalse(test swampopcodetype.Register, jump *swampopcodetype.Label) *BranchFalse {
	return &BranchFalse{test: test, jump: jump}
}

func (c *BranchFalse) Write(writer OpcodeWriter) error {
	writer.Command(CmdBranchFalse)
	writer.Register(c.test)
	writer.Label(c.jump)

	return nil
}

func (c *BranchFalse) String() string {
	return fmt.Sprintf("brfa %v %v", c.test, c.jump)
}
