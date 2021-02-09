/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcode

import (
	swampopcodetype "github.com/swamp/opcodes/type"
)

type LabelInject struct {
	referencedLabel      *swampopcodetype.Label
	logicalOrigoPosition swampopcodetype.ProgramCounter
	positionInStream     swampopcodetype.ProgramCounter
	offsetLabel          *swampopcodetype.Label
}

func NewLabelInject(l *swampopcodetype.Label, positionInStream swampopcodetype.ProgramCounter) *LabelInject {
	return &LabelInject{referencedLabel: l, positionInStream: positionInStream,
		logicalOrigoPosition: positionInStream.Add(1)}
}

func NewLabelInjectWithOffset(l *swampopcodetype.Label, positionInStream swampopcodetype.ProgramCounter,
	offsetLabel *swampopcodetype.Label) *LabelInject {
	return &LabelInject{referencedLabel: l, positionInStream: positionInStream,
		logicalOrigoPosition: positionInStream.Add(1), offsetLabel: offsetLabel}
}

func (l *LabelInject) ReferencedLabel() *swampopcodetype.Label {
	return l.referencedLabel
}

func (l *LabelInject) ForwardDeltaPC() (swampopcodetype.DeltaPC, error) {
	targetPc := l.referencedLabel.DefinedProgramCounter()
	if l.offsetLabel != nil {
		beforePc := l.offsetLabel.DefinedProgramCounter()
		newPc, newPcErr := DeltaProgramCounter(targetPc, beforePc)
		return newPc, newPcErr
	}
	return DeltaProgramCounter(targetPc, l.logicalOrigoPosition)
}

func (l *LabelInject) LocatedAtPosition() swampopcodetype.ProgramCounter {
	return l.positionInStream
}
