package nodes

import "testing"


/**
	Ensure that if we get a key, and removes a node, and re-add that node, we get the same.
*/
func TestAddingNodes(t *testing.T) {
	node := Node{id: "node1"}
	nodeArr := []Node{node}
	var cluster = NodeCluster{nodes: nodeArr }

	if len(cluster.nodes) != 1 {
		t.Error("Incorrect size of node array while initializing it.")
	}

	node2 := Node{id: "node2"}
	cluster.addNode(&node2)

	if len(cluster.nodes) != 2 {
		t.Error("Incorrect size after adding new node to cluster")
	}

}

func TestRemovingNodes(t *testing.T) {
	node := Node{id: "node1"}
	nodeArr := []Node{node}
	var cluster = NodeCluster{nodes: nodeArr }

	if len(cluster.nodes) != 1 {
		t.Error("Incorrect size of node array while initializing it.")
	}

	node2 := Node{id: "node2"}
	// Add node 2
	cluster.addNode(&node2)

	if len(cluster.nodes) != 2 {
		t.Error("Incorrect size after adding new node to cluster")
	}
	// Remove node 2
	cluster.removeNode(&node2)

	if len(cluster.nodes) != 1 {
		t.Error("Incorrect size after adding new node to cluster")
	}

}

