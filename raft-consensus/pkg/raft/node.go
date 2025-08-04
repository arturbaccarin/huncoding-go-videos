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
	log.Printf("Node %s set currentTerm local var", n.config.ID)

	if len(n.log) == 0 {
		log.Printf("Node %s WARNING: log is empty", n.config.ID)
	}

	log.Printf("Node %s about to unlock mutex", n.config.ID)
	n.mu.Unlock()
	log.Printf("Node %s unlocked mutex", n.config.ID)

	log.Printf("Node %s starting election for term %d", n.config.ID, currentTerm)

	votes := 1
	votesCh := make(chan bool, len(n.config.Peers))
	responses := 0
	responsesCh := make(chan struct{}, len(n.config.Peers))

	for _, peer := range n.config.Peers {
		log.Printf("Node %s sending RequestVote to %s", n.config.ID, peer)
		go func(peer string) {
			args := &RequestVoteArgs{
				Term:         currentTerm,
				Candidate:    n.config.ID,
				LastLogIndex: len(n.log) - 1,
				LastLogTerm:  n.log[len(n.log)-1].Term,
			}

			reply := &RequestVoteReply{}

			if err := n.sendRequestVote(peer, args, reply); err != nil {
				n.mu.Lock()
				defer n.mu.Unlock()

				if n.state != Candidate || n.currentTerm != currentTerm {
					return
				}

				responses++
				responsesCh <- struct{}{}

				if reply.VoteGranted {
					log.Printf("Node %s received vote from %s for term %d", n.config.ID, peer, currentTerm)
					votesCh <- true
				} else if reply.Term > currentTerm {
					log.Printf("Node %s received higher term %d from %s, updating state", n.config.ID, reply.Term, peer)
					n.becomeFollower(reply.Term)
				} else {
					log.Printf("Node %s received negative vote from %s for term %d", n.config.ID, peer, currentTerm)
				}
			} else {
				log.Printf("Node %s failed to send RequestVote to %s: %v", n.config.ID, peer, err)
				responses++
				responsesCh <- struct{}{}
			}
		}(peer)
	}
}
