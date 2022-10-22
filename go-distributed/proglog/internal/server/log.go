package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
  return &Log{}
}

func (self *Log) Append(record Record) (uint64, error) {
	self.lock()
	defer self.unlock()

	record.Offset = uint64(len(self.records))
	self.records = append(self.records, record)

	return record.Offset, nil
}

func (self *Log) Read(offset uint64) (Record, error) {
	self.lock()
	defer self.unlock()

	if offset >= uint64(len(self.records)) {
		return Record{}, ErrorOffsetNotFound
	}

	return self.records[offset], nil
}

func (self *Log) lock() {
	self.mu.Lock()
}

func (self *Log) unlock() {
	self.mu.Unlock()
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrorOffsetNotFound = fmt.Errorf("offset not found")
