package raft

import "time"

type AppendEntriesReply struct {
	Term    int
	Success bool
}

type Config struct {
	ID               string
	Peers            []string
	ElectionTimeout  time.Duration
	HeartbeatTimeout time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		ElectionTimeout:  150 * time.Millisecond,
		HeartbeatTimeout: 50 * time.Millisecond,
	}
}
