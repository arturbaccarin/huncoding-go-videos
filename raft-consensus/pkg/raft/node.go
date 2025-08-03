package raft

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

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
	if config == nil {
		config = DefaultConfig()
	}

	node := &Node{
		config:     config,
		state:      Follower,
		nextIndex:  make(map[string]int),
		matchIndex: make(map[string]int),
		applyCh:    make(chan interface{}, 100),
		stopCh:     make(chan struct{}),
	}

	node.log = append(node.log, LogEntry{Term: 0})

	return node
}

func (n *Node) Start() {
	go n.runElectionTimer()
}

func (n *Node) Stop() {
	close(n.stopCh)
}

func (n *Node) runElectionTimer() {
	for {
		timeout := n.config.ElectionTimeout + time.Duration(rand.Int63n(int64(n.config.ElectionTimeout)))
		log.Printf("Node %s waiting for %v before next election check", n.config.ID, timeout)

		select {
		case <-time.After(timeout):
			log.Printf("Node %s about to acquire lock in runElectionTimer", n.config.ID)
			n.mu.Lock()

			log.Printf("Node %s acquired lock in runElectionTimer", n.config.ID)
			if n.state != Leader {
				log.Printf("Node %s election timeout, current state: %v", n.config.ID, n.state)
				go n.startElection()
			}

			log.Printf("Node %s releasing lock in runElectionTimer", n.config.ID)
			n.mu.Unlock()
		case <-n.stopCh:
			return
		}
	}
}

func (n *Node) startElection() {
	log.Printf("Node %s ENTERED startElection", n.config.ID)

	n.mu.Lock()
	log.Printf("Node %s acquired lock in startElection", n.config.ID)

	if n.state == Leader {
		log.Printf("Node %s is already a leader, skipping election", n.config.ID)
		n.mu.Unlock()
		log.Printf("Node %s exiting startElection", n.config.ID)
		return
	}

	log.Printf("Node %s about to set state to Candidate", n.config.ID)
	n.state = Candidate
	log.Printf("Node %s set state to Candidate", n.config.ID)

	n.currentTerm++
	log.Printf("Node %s incremented current term to %d", n.config.ID, n.currentTerm)

	n.votedFor = n.config.ID
	log.Printf("Node %s set votedFor to %d", n.config.ID, n)

	currentTerm := n.currentTerm
	log.Printf()
}
