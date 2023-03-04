package server

import (
	"fmt"
	"sync"

	log_v1 "github.com/olee12/proglog/api/v1"
)

var ErrNotFound = fmt.Errorf("offset not found")

type Log struct {
	mu      sync.Mutex
	records []log_v1.Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record log_v1.Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (log_v1.Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return log_v1.Record{}, ErrNotFound
	}
	return c.records[offset], nil
}
