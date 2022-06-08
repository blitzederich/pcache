// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package binlog

import "errors"

var (
	ErrIncompleteRecord = errors.New("INCOMPLETE_RECORD")
)
