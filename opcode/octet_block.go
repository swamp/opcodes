/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcode

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type OctetBlock struct {
	octets []uint8
}

func NewOctetBlock(octets []uint8) *OctetBlock {
	return &OctetBlock{octets: octets}
}

func (o *OctetBlock) Octets() []uint8 {
	return o.octets
}

func (o *OctetBlock) Replace(pos swampopcodetype.ProgramCounter, v swampopcodetype.DeltaPC) {
	o.octets[pos.Value()] = uint8(v)
}

func (o *OctetBlock) FixUpLabelInject(r *LabelInject) error {
	referencedLabel := r.ReferencedLabel()
	if !referencedLabel.IsDefined() {
		return fmt.Errorf("label %v is not defined", referencedLabel)
	}

	delta, deltaErr := r.ForwardDeltaPC()
	if deltaErr != nil {
		return deltaErr
	}

	o.Replace(r.LocatedAtPosition(), delta)

	return nil
}

func (o *OctetBlock) FixUpLabelInjects(injects []*LabelInject) error {
	for _, inject := range injects {
		fixupErr := o.FixUpLabelInject(inject)
		if fixupErr != nil {
			return fixupErr
		}
	}

	return nil
}
