package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func NodeStatus(nodeList []*Node) {
	fmt.Println()
	for _, node := range nodeList {
		cpuAvailable, memoryAvailable := node.GetNodeStatus()
		fmt.Printf("Node %v has %v CPU and %v Memory available \n", node.id, cpuAvailable, memoryAvailable)
	}
}

func ReleaseNode(nodeList []*Node, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(5 * time.Second)
		randomNode := rand.Intn(10)
		nodeList[randomNode].Refresh()
	}

}

func AllotNodeToRequest(request *Request, nodeList []*Node) (*Node, error) {
	for _, node := range nodeList {
		if node.CanAllot(request) {
			return node, nil
		}
	}
	return nil, fmt.Errorf("there is no node which has enough resources\n")
}

func HandleRequest(request *Request, nodeList []*Node, wg *sync.WaitGroup) {
	defer wg.Done()
	node, err := AllotNodeToRequest(request, nodeList)
	if err != nil {
		fmt.Printf("Unable to allot node %s", err.Error())
	} else {
		fmt.Printf("Alloted node %v to request %v\n", node.id, request.id)
	}

}

func PrintNodeStatusPeriodically(nodeList []*Node, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		NodeStatus(nodeList)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	fmt.Println("Creating a system to manage resources")

	fmt.Println("Creating resource controller with few nodes")

	nodeList := []*Node{}
	for i := 0; i < 10; i++ {
		node := NewNode(i)
		nodeList = append(nodeList, node)
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		cpuRequired := rand.Intn(10)
		memoryRequired := rand.Intn(10)
		request := NewRequest(cpuRequired, memoryRequired, i)
		wg.Add(1)
		go HandleRequest(request, nodeList, wg)

	}

	wg.Add(1)
	go PrintNodeStatusPeriodically(nodeList, wg)

	wg.Add(1)
	go ReleaseNode(nodeList, wg)

	wg.Wait()

}
