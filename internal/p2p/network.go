package p2p

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

// Message represents a P2P message
type Message struct {
	From    string      `json:"from"`
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// NetNode represents a functional network node
type NetNode struct {
	ID        string
	Address   string
	Peers     []string
	mu        sync.RWMutex
	Consensus *PBFTNet
}

type PBFTNet struct {
	NodeID   string
	Node     *NetNode
	commits  map[string]int
	mu       sync.Mutex
}

func NewNetNode(id, addr string, peers []string) *NetNode {
	n := &NetNode{
		ID:      id,
		Address: addr,
		Peers:   peers,
	}
	n.Consensus = &PBFTNet{
		NodeID:  id,
		Node:    n,
		commits: make(map[string]int),
	}
	return n
}

func (n *NetNode) Start() {
	ln, err := net.Listen("tcp", n.Address)
	if err != nil {
		fmt.Printf("Node %s failed to listen: %v\n", n.ID, err)
		return
	}
	fmt.Printf("Node %s listening on %s\n", n.ID, n.Address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go n.handleConnection(conn)
	}
}

func (n *NetNode) handleConnection(conn net.Conn) {
	defer conn.Close()
	var msg Message
	if err := json.NewDecoder(conn).Decode(&msg); err != nil {
		return
	}
	fmt.Printf("Node %s: Received %s from %s\n", n.ID, msg.Type, msg.From)
	n.Consensus.Process(msg)
}

func (n *NetNode) Broadcast(msg Message) {
	for _, peer := range n.Peers {
		go n.send(peer, msg)
	}
}

func (n *NetNode) send(addr string, msg Message) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	defer conn.Close()
	json.NewEncoder(conn).Encode(msg)
}

func (p *PBFTNet) Process(msg Message) {
	p.mu.Lock()
	defer p.mu.Unlock()

	switch msg.Type {
	case "PROPOSAL":
		fmt.Printf("Node %s: Validating proposal...\n", p.NodeID)
		p.Node.Broadcast(Message{From: p.NodeID, Type: "COMMIT", Payload: msg.Payload})
	case "COMMIT":
		blockID := fmt.Sprintf("%v", msg.Payload)
		p.commits[blockID]++
		if p.commits[blockID] >= 2 {
			fmt.Printf("Node %s: BLOCK %s COMMITTED TO LEDGER\n", p.NodeID, blockID)
		}
	}
}
