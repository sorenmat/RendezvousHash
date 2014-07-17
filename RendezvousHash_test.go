package hashing

import "testing"


func TestReAddingNodes(t *testing.T) {
	var nodes []string
	nodes = addNode("node1", nodes)
	nodes = addNode("node2", nodes)
	var myNode = get("key", nodes)

	nodes = removeNode(myNode, nodes)
	nodes = addNode(myNode, nodes)

	if myNode != get("key", nodes) {
		t.Error("After we re-added a node, they are different ", myNode)
	}
}

func TestGetAndEmptyNode(t *testing.T) {
	var nodes[] string
	if get("key", nodes) != "" {
		t.Error("Got an non empty node")
	}
}
