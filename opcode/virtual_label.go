/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcode

import (
	"fmt"

	swampopcodeinst "github.com/swamp/opcodes/instruction"
	swampopcodetype "github.com/swamp/opcodes/type"
)

type VirtualLabel struct {
	label *swampopcodetype.Label
}

func NewVirtualLabel(label *swampopcodetype.Label) *VirtualLabel {
	return &VirtualLabel{label: label}
}

func (c *VirtualLabel) Label() *swampopcodetype.Label {
	return c.label
}

func (c *VirtualLabel) Write(writer swampopcodeinst.OpcodeWriter) error {
	return nil
}

func (c *VirtualLabel) String() string {
	return fmt.Sprintf("label: %v", c.label)
}
