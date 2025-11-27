package domain

type ClusterConfig struct {
	ReplicationFactor int
	ReadQuorum        int
	WriteQuorum       int
	VirtualNodes      int
}

// creates default cluster config.

func NewClusterConfig() *ClusterConfig {
	return &ClusterConfig{
		ReplicationFactor: 3,
		ReadQuorum:        2,
		WriteQuorum:       2,
		VirtualNodes:      150,
	}
}

func (c *ClusterConfig) Validate() error {
	if c.ReadQuorum > c.ReplicationFactor {
		return ErrInvalidQuorum
	}
	if c.WriteQuorum > c.ReplicationFactor {
		return ErrInvalidQuorum
	}
	if c.ReadQuorum+c.WriteQuorum <= c.ReplicationFactor {
		return ErrWeakConsistency
	}
	return nil
}
