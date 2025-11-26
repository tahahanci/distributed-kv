package domain

import (
	"fmt"
	"time"
)

type NodeStatus string

const (
	NodeStatusActive   NodeStatus = "active"
	NodeStatusInactive NodeStatus = "inactive"
	NodeStatusFailed   NodeStatus = "failed"
)

type Node struct {
	ID            string
	Address       string
	Port          int
	Status        NodeStatus
	LastHeartbeat time.Time
	HashValue     uint32
}

func NewNode(id, address string, port int) *Node {
	return &Node{
		ID:            id,
		Address:       address,
		Port:          port,
		Status:        NodeStatusActive,
		LastHeartbeat: time.Now(),
	}
}

func (n *Node) GetFullAddress() string {
	return fmt.Sprintf("%s:%d", n.Address, n.Port)
}

func (n *Node) IsHealthy(timeout time.Duration) bool {
	return n.Status == NodeStatusActive &&
		time.Since(n.LastHeartbeat) < timeout
}

func (n *Node) UpdateHeartbeat() {
	n.LastHeartbeat = time.Now()
	n.Status = NodeStatusActive
}
