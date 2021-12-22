package opcode_sp_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/swamp/opcodes/opcode_sp"
)

func TestSomething(t *testing.T) {
	debugInfo := opcode_sp.NewDebugInfo()

	filePosition := opcode_sp.FilePosition{
		SourceFileID: 42,
		Line:         3,
		Column:       99,
	}

	debugInfo.AddOpcodeSource(20, filePosition)

	octets, err := debugInfo.Octets()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.Dump(octets))
}
