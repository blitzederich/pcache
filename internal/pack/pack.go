// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package pack

import "bytes"

const (
	B_0 = 0xFF
	B_1 = 0xFF00
	B_2 = 0xFF0000
	B_3 = 0xFF000000

	B_4 = 0xFF00000000
	B_5 = 0xFF0000000000
	B_6 = 0xFF000000000000
	B_7 = 0xFF00000000000000
)

// convert uint32 to LittleEndian bytes
func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n & B_0),
		byte(n & B_1 >> 8),
		byte(n & B_2 >> 16),
		byte(n & B_3 >> 24),
	}
}

// convert LittleEndian bytes to uint32
func BytesToUint32(b []byte) uint32 {
	n := uint32(b[3])
	n = n<<8 + uint32(b[2])
	n = n<<8 + uint32(b[1])
	n = n<<8 + uint32(b[0])
	return n
}

// convert uint64 to LittleEndian bytes
func Uint64ToBytes(n uint64) []byte {
	return []byte{
		byte(n & B_0),
		byte(n & B_1 >> 8),
		byte(n & B_2 >> 16),
		byte(n & B_3 >> 24),
		byte(n & B_4 >> 32),
		byte(n & B_5 >> 40),
		byte(n & B_6 >> 48),
		byte(n & B_7 >> 56),
	}
}

// convert LittleEndian bytes to uint64
func BytesToUint64(b []byte) uint64 {
	n := uint64(b[7])
	n = n<<8 + uint64(b[6])
	n = n<<8 + uint64(b[5])
	n = n<<8 + uint64(b[4])
	n = n<<8 + uint64(b[3])
	n = n<<8 + uint64(b[2])
	n = n<<8 + uint64(b[1])
	n = n<<8 + uint64(b[0])
	return n
}

// convert string to bytes with one byte prefix
func PackShortString(s string) []byte {
	return bytes.Join([][]byte{
		{uint8(len(s))},
		[]byte(s),
	}, []byte{})
}

// convert string to bytes with 4 LittleEndian bytes prefix
func PackString(s string) []byte {
	return bytes.Join([][]byte{
		Uint32ToBytes(uint32(len(s))),
		[]byte(s),
	}, []byte{})
}

// convert string to bytes with 8 LittleEndian bytes prefix
func PackLongString(s string) []byte {
	return bytes.Join([][]byte{
		Uint64ToBytes(uint64(len(s))),
		[]byte(s),
	}, []byte{})
}
