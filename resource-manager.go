package main

import "sync"

type Node struct {
	memoryAvailable int
	cpuAvailable    int
	requestList     []Request
	id              int
	mutex           sync.RWMutex
}

func NewNode(id int) *Node {
	node := Node{
		memoryAvailable: 10,
		cpuAvailable:    10,
		requestList:     []Request{},
		id:              id,
	}
	return &node
}

func (node *Node) CanAllot(request *Request) bool {
	node.mutex.Lock()
	defer node.mutex.Unlock()
	if request.memory <= node.memoryAvailable && request.cpu <= node.cpuAvailable {
		node.cpuAvailable -= request.cpu
		node.memoryAvailable -= request.memory
		request.nodeId = node.id
		node.requestList = append(node.requestList, *request)
		return true
	} else {
		return false
	}
}

func (node *Node) GetNodeStatus() (int, int) {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	cpu, mem := node.cpuAvailable, node.memoryAvailable
	return cpu, mem
}

func (node *Node) Refresh() {
	node.mutex.Lock()
	defer node.mutex.Unlock()
	node.cpuAvailable = 10
	node.memoryAvailable = 10
	node.requestList = node.requestList[:0]
}
