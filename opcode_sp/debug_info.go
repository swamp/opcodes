package opcode_sp

import (
	"encoding/binary"
	"fmt"
)

type OpcodePosition uint

type LineCol struct {
	Line uint
	Column uint
}

type FilePosition struct {
	SourceFileID uint
	Start LineCol
	End LineCol
}

type OpcodeInfo struct {
	filePosition   FilePosition
	opcodePosition OpcodePosition
}

type DebugInfo struct {
	lastOpcodePosition int
	infos              []OpcodeInfo
}

func NewDebugInfo() *DebugInfo {
	return &DebugInfo{
		lastOpcodePosition: -1,
	}
}

func (d *DebugInfo) AddOpcodeSource(position OpcodePosition, filePosition FilePosition) error {
	if int(position) < d.lastOpcodePosition {
		return fmt.Errorf("must have opcodes in ascending order")
	}

	if position > 0xf00 {
		panic("too big bvalue")
	}

	info := OpcodeInfo{filePosition: filePosition, opcodePosition: position}
	d.infos = append(d.infos, info)

	d.lastOpcodePosition = int(position)

	return nil
}


type octetStream struct {
	octets []byte
}

func (d *octetStream) writeUint16(v uint16) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	d.octets = append(d.octets, b[:]...)
}

func SerializeDebugLines(infos []OpcodeInfo) ([]byte, error) {
	stream := &octetStream{}
	lastPosition := -1
	for _, info := range infos {
		if int(info.opcodePosition) <= lastPosition {
			panic("opcode positions in wrong order")
		}
		if info.opcodePosition > 0xf00 {
			panic("too big bvalue")
		}
		lastPosition = int(info.opcodePosition)
		stream.writeUint16(uint16(info.opcodePosition))
		stream.writeUint16(uint16(info.filePosition.SourceFileID))
		stream.writeUint16(uint16(info.filePosition.Start.Line))
		stream.writeUint16(uint16(info.filePosition.Start.Column))
	}

	return stream.octets, nil
}
