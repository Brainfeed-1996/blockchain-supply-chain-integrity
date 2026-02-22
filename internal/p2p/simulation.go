package p2p

import (
	"fmt"
	"sync"
	"time"
)

// Message represents a P2P message
type Message struct {
	From    string
	Type    string
	Payload interface{}
}

// Node represents a simulated network node
type Node struct {
	ID           string
	Peers        map[string]*Node
	Inbox        chan Message
	ledger       chan interface{}
	mu           sync.RWMutex
	consensus    *PBFT
	StateMachine *StateMachine
}

func NewNode(id string) *Node {
	n := &Node{
		ID:           id,
		Peers:        make(map[string]*Node),
		Inbox:        make(chan Message, 100),
		StateMachine: NewStateMachine(),
	}
	n.consensus = NewPBFT(id, n)
	return n
}

func (n *Node) AddPeer(peer *Node) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Peers[peer.ID] = peer
}

func (n *Node) Broadcast(msgType string, payload interface{}) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	msg := Message{From: n.ID, Type: msgType, Payload: payload}
	for _, peer := range n.Peers {
		peer.Inbox <- msg
	}
}

func (n *Node) Start() {
	go func() {
		for msg := range n.Inbox {
			n.handleMessage(msg)
		}
	}()
}

func (n *Node) handleMessage(msg Message) {
	switch msg.Type {
	case "PRE-PREPARE":
		n.consensus.HandlePrePrepare(msg)
	case "PREPARE":
		n.consensus.HandlePrepare(msg)
	case "COMMIT":
		n.consensus.HandleCommit(msg)
	}
}

// PBFT State Machine Simulation
type PBFT struct {
	NodeID   string
	node     *Node
	sequence int
	view     int
	prepares map[int]map[string]bool
	commits  map[int]map[string]bool
	mu       sync.Mutex
}

func NewPBFT(id string, n *Node) *PBFT {
	return &PBFT{
		NodeID:   id,
		node:     n,
		prepares: make(map[int]map[string]bool),
		commits:  make(map[int]map[string]bool),
	}
}

func (p *PBFT) Propose(block interface{}) {
	p.sequence++
	fmt.Printf("Node %s: Proposing block %d\n", p.NodeID, p.sequence)
	p.node.Broadcast("PRE-PREPARE", block)
}

func (p *PBFT) HandlePrePrepare(msg Message) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Printf("Node %s: Received PRE-PREPARE from %s\n", p.NodeID, msg.From)
	// Validate and then prepare
	p.node.Broadcast("PREPARE", msg.Payload)
}

func (p *PBFT) HandlePrepare(msg Message) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	seq := 1 // Simplified
	if p.prepares[seq] == nil {
		p.prepares[seq] = make(map[string]bool)
	}
	p.prepares[seq][msg.From] = true
	
	if len(p.prepares[seq]) >= 2 { // Simplified threshold
		fmt.Printf("Node %s: Reached PREPARE quorum for %d\n", p.NodeID, seq)
		p.node.Broadcast("COMMIT", msg.Payload)
	}
}

func (p *PBFT) HandleCommit(msg Message) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	seq := 1 // Simplified
	if p.commits[seq] == nil {
		p.commits[seq] = make(map[string]bool)
	}
	p.commits[seq][msg.From] = true
	
	if len(p.commits[seq]) >= 2 { // Simplified threshold
		fmt.Printf("Node %s: Reached COMMIT quorum for %d. Block Validated!\n", p.NodeID, seq)
	}
}
