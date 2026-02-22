package main

import (
	"blockchain-supply-chain-integrity/internal/p2p"
	"os"
	"strings"
	"time"
)

func main() {
	nodeID := os.Getenv("NODE_ID")
	addr := os.Getenv("NODE_ADDR")
	peersStr := os.Getenv("PEERS")
	
	peers := []string{}
	if peersStr != "" {
		peers = strings.Split(peersStr, ",")
	}

	node := p2p.NewNetNode(nodeID, addr, peers)
	
	if nodeID == "node1" {
		go func() {
			time.Sleep(10 * time.Second)
			node.Broadcast(p2p.Message{
				From:    nodeID,
				Type:    "PROPOSAL",
				Payload: "Batch-2026-XYZ",
			})
		}()
	}

	node.Start()
}
