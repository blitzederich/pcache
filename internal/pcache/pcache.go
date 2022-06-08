// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package pcache

import (
	"fmt"
	"log"
	"os"
	"pcache/internal/binlog"
	"pcache/internal/store"
)

type Log interface {
	Set(key, value string) error
	Update(key, value string) error
	Delete(key string) error

	Parse(cache binlog.Cache)
}

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string)
	Update(key, value string)
	Delete(key string)

	GetStore() *map[string]string
}

type PCache struct {
	logPath string
	log     Log
	cache   Cache
}

func New() PCache {
	return PCache{
		log:   nil,
		cache: nil,
	}
}

func (pc *PCache) Init(binLogPath string) error {
	pc.logPath = binLogPath
	file, err := os.OpenFile(pc.logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	pc.log = binlog.New(file)
	pc.cache = store.New[string, string]()
	pc.log.Parse(pc.cache)

	return nil
}

func (pc *PCache) Get(key string) (string, bool) {
	return pc.cache.Get(key)
}

func (pc *PCache) Set(key, value string) error {
	if err := pc.log.Set(key, value); err != nil {
		return err
	}

	pc.cache.Set(key, value)
	return nil
}

func (pc *PCache) Update(key, value string) error {
	if err := pc.log.Update(key, value); err != nil {
		return err
	}

	pc.cache.Update(key, value)
	return nil
}

func (pc *PCache) Delete(key string) error {
	if err := pc.log.Delete(key); err != nil {
		return err
	}

	pc.cache.Delete(key)
	return nil
}

func (pc *PCache) Stats() {
	store := pc.cache.GetStore()
	for k, v := range *store {
		fmt.Println(k, "\t", v)
	}
}

func (pc *PCache) Optimize() {
	file, err := os.OpenFile(pc.logPath+"_temp", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	bl := binlog.New(file)

	store := pc.cache.GetStore()
	for k, v := range *store {
		bl.Set(k, v)
	}

	if err := os.Remove(pc.logPath); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := os.Rename(pc.logPath+"_temp", pc.logPath); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	pc.Init(pc.logPath)
}
