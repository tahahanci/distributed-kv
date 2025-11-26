package domain

import "time"

type Entry struct {
	Key       string
	Value     []byte
	Timestamp time.Time
	Version   int64
	NodeID    string
}

func NewEntry(key string, value []byte, nodeID string) *Entry {
	return &Entry{
		Key:       key,
		Value:     value,
		NodeID:    nodeID,
		Timestamp: time.Now(),
		Version:   1,
	}
}

func (e *Entry) IsNewer(other *Entry) bool {
	if e.Version != other.Version {
		return e.Version > other.Version
	}
	return e.Timestamp.After(other.Timestamp)
}
