package raft

import "sync"

type Node struct {
	mu     sync.Mutex
	config *Config

	// Persistent state
	currentTerm int
	votedFor    string
	log         []LogEntry

	// Volatile state
	state       NodeState
	commitIndex int
	lastApplied int

	// Leader state
	nextIndex  map[string]int
	matchIndex map[string]int

	// Channels
	applyCh chan interface{}
	stopCh  chan struct{}
}

func NewNode(config *Config) *Node {

}
