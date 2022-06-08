// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package buffer

import (
	"errors"
	"io"
)

type Buffer struct {
	reader io.Reader
	buff   []byte
	index  int32
	len    int32
	end    bool
}

func New(r io.Reader, buffSize int) *Buffer {
	return &Buffer{
		reader: r,
		buff:   make([]byte, buffSize),
		index:  0,
		len:    0,
		end:    false,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.len == b.index {
		if err := b.readChunk(); err != nil {
			return 0, err
		}
	}

	i := b.index
	b.index++
	return b.buff[i], nil
}

func (b *Buffer) readChunk() error {
	if b.end {
		return io.EOF
	}

	n, err := b.reader.Read(b.buff)
	if err != nil {
		if errors.Is(err, io.EOF) {
			b.end = true
		}
		return err
	}

	b.len = int32(n)
	b.index = 0
	return nil
}
