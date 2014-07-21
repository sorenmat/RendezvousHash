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

/**
	Find the node that is responsible for the key.
	This is done by iterating over all the known nodes, and hashing the key with the id on the current node. The combination
	if the hashed key + hashed id with the highest number is the node we choose.
*/
func (cluster *NodeCluster) whichNodeToPutKey(key string) string {
	orderedNodes := cluster.nodes
	var maxValue uint64 = 0
	var max string
	for i := 0; i < len(orderedNodes); i++ {
		var nodeHash = toHash(key, orderedNodes[i].id)
		if nodeHash > maxValue {
			max = orderedNodes[i].id
			maxValue = nodeHash
		}
	}
	return max
}

func (cluster *NodeCluster) getNodes() []Node {
	return cluster.nodes
}