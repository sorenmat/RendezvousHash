package nodes

type Node struct {
	id string
}

type NodeCluster struct {
	nodes []Node
}

/**
	Add a new node to the cluster
*/
func (cluster *NodeCluster) addNode(node *Node) {
	var arr = cluster.nodes
	for i := 0; i < len(arr); i++ {
		if arr[i].id == node.id {
			panic("A node with that id exists in the cluster")	
		}
	}

	cluster.nodes = append(cluster.nodes, *node)
}

/**
	Remove a node from the cluster by comparing its id
*/
func (cluster *NodeCluster) removeNode(node *Node) {
	arr := cluster.nodes
	for i := 0; i < len(arr); i++ {
			if arr[i].id == node.id {
				arr = append(arr[:i], arr[i+1:]...)				
			}
	}
	cluster.nodes = arr
}


func (cluster *NodeCluster) getNodes() []Node {
	return cluster.nodes
}