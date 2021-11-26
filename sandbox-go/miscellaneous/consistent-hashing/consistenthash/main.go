package main

import (
	"fmt"

	chash "github.com/mbrostami/consistenthash"
)

func main() {
	// Create ConsistentHash with 2 replicas
	ch := chash.NewConsistentHash(2, nil)
	ch.Add("127.0.0.1:1001") // node 1
	ch.Add("127.0.0.1:1002") // node 2
	// ch.AddReplicas("127.0.0.1:1003", 4) // node3 has more capacity so possibility to get assigned request is higher than other nodes

	rKey := "something like request url"
	node := ch.Get(rKey) // find upper closest node
	fmt.Println(node)    // this will print out one of the nodes
}
