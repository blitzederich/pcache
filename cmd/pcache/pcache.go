// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package main

import (
	"flag"
	"log"
	"os"
	"pcache/internal/pcache"
)

var (
	argv struct {
		binLogPath string
	}
)

func main() {
	flag.StringVar(&argv.binLogPath, "binlog_path", ".binlog", "the path to binlog")
	flag.Parse()

	pcache := pcache.New()
	if err := pcache.Init(argv.binLogPath); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	pcache.Set("key_1", "value_1")
	pcache.Set("key_2", "value_2")
	pcache.Update("key_1", "new_value_2")
	pcache.Set("key_3", "value_3")
	pcache.Delete("key_2")

	// pcache.Optimize()

	pcache.Stats()

}
