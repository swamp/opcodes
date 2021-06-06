/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodeinst

import (
	swampopcodetype "github.com/swamp/opcodes/type"
)

type OpcodeWriter interface {
	Register(r swampopcodetype.Register)
	Field(f swampopcodetype.Field)
	DeltaPC(pc swampopcodetype.DeltaPC)
	Label(l *swampopcodetype.Label)
	LabelWithOffset(l *swampopcodetype.Label, offset *swampopcodetype.Label)
	EnumValue(v uint8)
	Count(c int)
	TypeIDConstant(c uint16)
	Command(cmd Commands)
}
