package main

import "fmt"

type node struct {
	label string
	value int
	edges []*node
}

func main() {
	node1 := &node{label: "node1", value: 1}
	node2 := &node{label: "node2", value: 2}
	node3 := &node{label: "node3", value: 3}
	node4 := &node{label: "node4", value: 4}

	node1.edges = append(node1.edges, node2)
	node1.edges = append(node1.edges, node3)
	node1.edges = append(node1.edges, node4)
	node4.edges = append(node4.edges, node2)

	fmt.Println(node1)

	for i := range node1.edges {
		fmt.Println(node1.edges[i].edges)
	}
}
