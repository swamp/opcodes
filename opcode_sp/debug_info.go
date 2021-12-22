package opcode_sp

import (
	"encoding/binary"
	"fmt"
	"io"
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
	octets             []byte
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

	info := OpcodeInfo{filePosition: filePosition, opcodePosition: position}
	d.infos = append(d.infos, info)

	d.lastOpcodePosition = int(position)

	return nil
}

func (d *DebugInfo) writeUint16(v uint16) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	d.octets = append(d.octets, b[:]...)
}

func (d *DebugInfo) Octets() ([]byte, error) {
	d.writeUint16(uint16(len(d.infos)))

	for _, info := range d.infos {
		d.writeUint16(uint16(info.opcodePosition))
		d.writeUint16(uint16(info.filePosition.SourceFileID))
		d.writeUint16(uint16(info.filePosition.Start.Line))
		d.writeUint16(uint16(info.filePosition.Start.Column))
	}

	return d.octets, nil
}

func (d *DebugInfo) Serialize(writer io.Writer) error {

	_, err := writer.Write(d.octets)
	return err
}
