// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package pack

import (
	"testing"
)

func TestUint32ToBytes(t *testing.T) {
	expect := []byte{231, 199, 176, 124}
	result := Uint32ToBytes(uint32(2091960295))

	if string(expect) != string(result) {
		t.Errorf("Expected: %v, Result: %v", expect, result)
	}
}

func TestBytesToUint32(t *testing.T) {
	expect := uint32(2091960295)
	result := BytesToUint32([]byte{231, 199, 176, 124})

	if expect != result {
		t.Errorf("Expected: %v, Result: %v", expect, result)
	}
}

func TestUint64ToBytes(t *testing.T) {
	expect := []byte{47, 170, 58, 118, 160, 110, 7, 0}
	result := Uint64ToBytes(uint64(2091960294353455))

	if string(expect) != string(result) {
		t.Errorf("Expected: %v, Result: %v", expect, result)
	}
}

func TestBytesToUint64(t *testing.T) {
	expect := uint64(2091960294353455)
	result := BytesToUint64([]byte{47, 170, 58, 118, 160, 110, 7, 0})

	if expect != result {
		t.Errorf("Expected: %v, Result: %v", expect, result)
	}
}

func TestPackShortString(t *testing.T) {
	expect := []byte{12, 104, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100}
	result := PackShortString("hello, world")

	if string(expect) != string(result) {
		t.Errorf("Expected: %v, Result: %v", expect, result)
	}
}

func TestPackString(t *testing.T) {
	expect := []byte{12, 0, 0, 0, 104, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100}
	result := PackString("hello, world")

	if string(expect) != string(result) {
		t.Errorf("Expected: %v, Result: %v", expect, result)
	}
}

func TestPackLongString(t *testing.T) {
	expect := []byte{12, 0, 0, 0, 0, 0, 0, 0, 104, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100}
	result := PackLongString("hello, world")

	if string(expect) != string(result) {
		t.Errorf("Expected: %v, Result: %v", expect, result)
	}
}
