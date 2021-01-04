/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcode

import (
	swampopcodeinst "github.com/swamp/opcodes/instruction"
	swampopcodetype "github.com/swamp/opcodes/type"
)

type OpCodeStream struct {
	octets       []byte
	labelInjects []*LabelInject
}

func NewOpCodeStream() *OpCodeStream {
	return &OpCodeStream{}
}

func (s *OpCodeStream) LabelInjects() []*LabelInject {
	return s.labelInjects
}

func (s *OpCodeStream) Octets() []byte {
	return s.octets
}

func (s *OpCodeStream) Write(c uint8) {
	s.octets = append(s.octets, c)
}

func (s *OpCodeStream) Register(r swampopcodetype.Register) {
	s.Write(r.Value())
}

func (s *OpCodeStream) Field(f swampopcodetype.Field) {
	s.Write(f.Value())
}

func (s *OpCodeStream) EnumValue(v uint8) {
	s.Write(v)
}

func (s *OpCodeStream) programCounter() swampopcodetype.ProgramCounter {
	return swampopcodetype.NewProgramCounter(uint16(len(s.octets)))
}

func (s *OpCodeStream) DeltaPC(pc swampopcodetype.DeltaPC) {
	s.Write(uint8(pc))
}

func (s *OpCodeStream) addLabelInject(inject *LabelInject) {
	s.DeltaPC(0xff)
	s.labelInjects = append(s.labelInjects, inject)
}

func (s *OpCodeStream) Label(l *swampopcodetype.Label) {
	inject := NewLabelInject(l, s.programCounter())
	s.addLabelInject(inject)
}

func (s *OpCodeStream) LabelWithOffset(l *swampopcodetype.Label, offset *swampopcodetype.Label) {
	inject := NewLabelInjectWithOffset(l, s.programCounter(), offset)
	s.addLabelInject(inject)
}

func (s *OpCodeStream) Count(c int) {
	s.Write(uint8(c))
}

func (s *OpCodeStream) Command(cmd swampopcodeinst.Commands) {
	s.Write(uint8(cmd))
}
