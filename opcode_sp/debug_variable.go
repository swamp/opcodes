package opcode_sp

func SerializeDebugScopes(infos []VariableInfo) ([]byte, error) {
	stream := &octetStream{}
	lastPosition := -1
	for _, info := range infos {
		if int(info.StartOpcodePosition) <= lastPosition {
			panic("opcode positions in wrong order")
		}
		if info.StartOpcodePosition > 0xf00 {
			panic("too big bvalue")
		}

		lastPosition = int(info.StartOpcodePosition)
		stream.writeUint16(uint16(info.StartOpcodePosition))
		stream.writeUint16(uint16(info.EndOpcodePosition))
		stream.writeUint16(uint16(info.TypeID))
		stream.writeUint16(uint16(info.ScopeID))
	}

	return stream.octets, nil
}

