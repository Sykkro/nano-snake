package state

func setBit(b byte, i uint8) byte {
	b |= (1 << i)
	return b
}

func clearBit(b byte, i uint8) byte {
	b &= ^(1 << i)
	return b
}

func hasBit(b byte, i uint8) bool {
	return b&(1<<i) > 0
}

func xyToByte(x uint8, y uint8) byte {
	return (x << 4) | y
}

func byteToXY(b byte) (x uint8, y uint8) {
	return b >> 4, b & 0x0F
}
