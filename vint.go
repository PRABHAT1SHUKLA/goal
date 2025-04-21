func ReadVarInt(conn net.Conn) (int, error) {
	var num int
	var numRead int
	var b [1]byte
	for {
		_, err := conn.Read(b[:])
		if err != nil {
			return 0, err
		}
		byteVal := b[0]
		num |= int(byteVal&0x7F) << (7 * numRead)
		numRead++
		if numRead > 5 {
			return 0, fmt.Errorf("VarInt is too big")
		}
		if byteVal&0x80 == 0 {
			break
		}
	}
	return num, nil
}
