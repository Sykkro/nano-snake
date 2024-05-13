package binary

func SetBit(b byte, i uint8) byte {
	b |= (1 << i)
	return b
}

func ClearBit(b byte, i uint8) byte {
	b &= ^(1 << i)
	return b
}

func HasBit(b byte, i uint8) bool {
	return b&(1<<i) > 0
}

func XyToByte(x uint8, y uint8) byte {
	return (x << 4) | y
}

func ByteToXY(b byte) (x uint8, y uint8) {
	return b >> 4, b & 0x0F
}
