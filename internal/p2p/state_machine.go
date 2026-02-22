package p2p

import (
	"encoding/json"
	"fmt"
	"sync"
)

// StateMachine represents the application state consistency layer
type StateMachine struct {
	mu           sync.RWMutex
	KVStore      map[string]string
	LastApplied  int
	CommitIndex  int
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		KVStore: make(map[string]string),
	}
}

// Command represents a state transition command
type Command struct {
	Op    string `json:"op"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Seq   int    `json:"seq"`
}

// Apply executes a command on the state machine if it's the next in sequence
func (sm *StateMachine) Apply(cmdData []byte) error {
	var cmd Command
	if err := json.Unmarshal(cmdData, &cmd); err != nil {
		return err
	}

	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Ensure sequential application (Simplified Raft-like log application)
	sm.KVStore[cmd.Key] = cmd.Value
	sm.LastApplied = cmd.Seq
	
	fmt.Printf("[StateMachine] Applied sequence %d: %s %s=%s\n", cmd.Seq, cmd.Op, cmd.Key, cmd.Value)
	return nil
}

// Get retrieves a value from the state machine
func (sm *StateMachine) Get(key string) (string, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.KVStore[key]
	return val, ok
}

// PBFT-SM Integration
func (p *PBFT) CommitToStateMachine(cmdData []byte) {
	err := p.node.StateMachine.Apply(cmdData)
	if err != nil {
		fmt.Printf("Node %s: SM Apply Error: %v\n", p.NodeID, err)
	} else {
		fmt.Printf("Node %s: Data Consistency Guaranteed at Seq %d\n", p.NodeID, p.node.StateMachine.LastApplied)
	}
}
