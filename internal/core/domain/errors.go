package domain

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyExists   = errors.New("key already exists")
	ErrStorageFull = errors.New("storage is full")

	ErrInvalidQuorum     = errors.New("invalid quorum configuration")
	ErrWeakConsistency   = errors.New("weak consistency: R + W <= N")
	ErrInsufficientNodes = errors.New("insufficient nodes for operation")
	ErrNodeNotFound      = errors.New("node not found")

	ErrReplicationFailed = errors.New("replication failed")
	ErrQuorumNotMet      = errors.New("quorum not met")
)
