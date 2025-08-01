package store

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/coreos/etcd/raft"
)

type Command struct {
	Op    string      `json:"op"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Store struct {
	mu    sync.RWMutex
	raft  *raft.Node
	store map[string]interface{}
}

func NewStore(raftNode *raft.Node) *Store {
	s := &Store{
		raft:  raftNode,
		store: make(map[string]interface{}),
	}

	go s.applyComittedEntries()

	return s
}

func (s *Store) Get(key string) (interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if value, ok := s.store[key]; ok {
		return value, nil
	}

	return nil, errors.New("key not found")
}

func (s *Store) Set(key string, value interface{}) error {
	cmd := Command{
		Op:    "SET",
		Key:   key,
		Value: value,
	}

	if _, err := json.Marshal(cmd); err != nil {
		return err
	}

	s.mu.Lock()
	s.store[key] = value
	s.mu.Unlock()

	return nil
}

func (s *Store) Delete(key string) error {
	cmd := Command{
		Op:  "DELETE",
		Key: key,
	}

	if _, err := json.Marshal(cmd); err != nil {
		return err
	}

	s.mu.Lock()
	delete(s.store, key)
	s.mu.Unlock()

	return nil
}

func (s *Store) applyComittedEntries() {
	for {
		select {
		case entry := <-s.raft.GetApplyCh():
			var cmd Command
			if err := json.Unmarshal(entry.([]byte), &cmd); err != nil {
				continue
			}

			s.mu.Lock()
			switch cmd.Op {
			case "SET":
				s.store[cmd.Key] = cmd.Value
			case "DELETE":
				delete(s.store, cmd.Key)
			}
			s.mu.Unlock()
		}
	}
}
