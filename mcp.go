package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func readVarInt(conn net.Conn) (int, error) {
	var num int
	var numRead int
	var read byte
	for {
		if numRead > 5 {
			return 0, fmt.Errorf("VarInt too big")
		}
		if _, err := conn.Read([]byte{read}); err != nil {
			return 0, err
		}
		val := int(read & 0x7F)
		num |= val << (7 * numRead)
		numRead++
		if read&0x80 == 0 {
			break
		}
	}
	return num, nil
}

func writeVarInt(buf *bytes.Buffer, value int) {
	for {
		temp := byte(value & 0x7F)
		value >>= 7
		if value != 0 {
			temp |= 0x80
		}
		buf.WriteByte(temp)
		if value == 0 {
			break
		}
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read handshake
	_, _ = readVarInt(conn)
	data := make([]byte, 1024)
	n, _ := conn.Read(data)
	if n == 0 {
		return
	}

	// Respond to status ping
	resp := `{"version":{"name
