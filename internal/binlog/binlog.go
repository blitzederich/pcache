// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package binlog

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"pcache/internal/buffer"
	"pcache/internal/pack"
)

const (
	BUFF_SIZE = 1024
)

const (
	ACTION_SET    = "S"
	ACTION_UPDATE = "U"
	ACTION_DELETE = "D"
)

const (
	READ_ACTION_SIZE = 0
	READ_ACTION      = 1
	READ_KEY_SIZE    = 2
	READ_KEY         = 3
	READ_VALUE_SIZE  = 4
	READ_VALUE       = 5
)

type Cache interface {
	Set(key, value string)
	Update(key, value string)
	Delete(key string)
}

type BinLog struct {
	file *os.File
}

func New(file *os.File) *BinLog {
	return &BinLog{
		file: file,
	}
}

func (bl *BinLog) Set(key, value string) error {
	return bl.write(ACTION_SET, key, value)
}

func (bl *BinLog) Update(key, value string) error {
	return bl.write(ACTION_UPDATE, key, value)
}

func (bl *BinLog) Delete(key string) error {
	return bl.write(ACTION_DELETE, key, "")
}

func (bl *BinLog) write(action, key, value string) error {
	record := bytes.Join([][]byte{
		pack.PackShortString(action),
		pack.PackShortString(key),
		pack.PackString(value),
	}, []byte{})

	n, err := bl.file.Write(record)
	if err != nil {
		return err
	}

	if n < len(record) {
		return ErrIncompleteRecord
	}

	return nil
}

func (bl *BinLog) Parse(cache Cache) {
	buff := buffer.New(bl.file, BUFF_SIZE)

	var (
		actionSize uint8
		action     string
		keySize    uint8
		key        string
		valueSize  uint32
		value      string
	)

	var (
		counter uint32 = 0
		tmpBuff        = make([]byte, 0)
		state          = READ_ACTION_SIZE
	)

	for {
		c, err := buff.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
			os.Exit(1)
		}

	Switch:
		switch state {
		case READ_ACTION_SIZE:
			actionSize = c
			state = READ_ACTION
		case READ_ACTION:
			if counter == uint32(actionSize) {
				state = READ_KEY_SIZE
				counter = 0
				goto Switch
			}
			action += string(c)
			counter++
		case READ_KEY_SIZE:
			keySize = c
			state = READ_KEY
		case READ_KEY:
			if counter == uint32(keySize) {
				state = READ_VALUE_SIZE
				counter = 0
				goto Switch
			}
			key += string(c)
			counter++
		case READ_VALUE_SIZE:
			if counter == 4 {
				valueSize = pack.BytesToUint32(tmpBuff)
				tmpBuff = tmpBuff[:0]

				state = READ_VALUE
				counter = 0
				goto Switch
			}
			tmpBuff = append(tmpBuff, c)
			counter++
		case READ_VALUE:
			if valueSize != 0 {
				value += string(c)
				counter++
			}

			if counter == valueSize {
				switch action {
				case ACTION_SET:
					cache.Set(key, value)
				case ACTION_UPDATE:
					cache.Update(key, value)
				case ACTION_DELETE:
					cache.Delete(key)
				}

				actionSize, keySize, valueSize = 0, 0, 0
				action, key, value = "", "", ""

				state = READ_ACTION_SIZE
				counter = 0
			}
		}
	}
}
