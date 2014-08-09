package main

import "testing"
import "fmt"

/**
Ensure that if we get a key, and removes a node, and re-add that node, we get the same.
*/
func TestReAddingNodes(t *testing.T) {
	var nodes []string
	nodes = addNode("node1", nodes)
	nodes = addNode("node2", nodes)
	var myNode = whichNodeOwnsKey("key", nodes)

	nodes = removeNode(myNode, nodes)
	nodes = addNode(myNode, nodes)

	if myNode != whichNodeOwnsKey("key", nodes) {
		t.Error("After we re-added a node, they are different ", myNode)
	}
}

/**
Getting a key from an empty node list, should fail
*/
func TestGetAndEmptyNode(t *testing.T) {
	var nodes []string
	if whichNodeOwnsKey("key", nodes) != "" {
		t.Error("Got an non empty node")
	}

}

func TestGetNode(t *testing.T) {
	var nodes []string
	nodes = addNode("node1", nodes)
	nodes = addNode("node2", nodes)

	var myNode = whichNodeOwnsKey("key", nodes)

	fmt.Println(myNode)

}
