package hashing

import (
	"github.com/spaolacci/murmur3"
	"fmt"
	"sort"
)


func toHash(value ...string) uint64 {
	hasher := murmur3.New64()
	for _, v := range value {
    	hasher.Write([]byte(v))
    }	
	var hash = hasher.Sum64()
	return hash
}

func removeFromArray(item string, arr []string) []string {
	for i := 0; i < len(arr); i++ {
			if arr[i] == item {
				arr = append(arr[:i], arr[i+1:]...)				
			}
	}
	return arr	
}

func removeNode(node string, nodes []string) []string {
		var newnodes = removeFromArray(node, nodes)
		sort.Strings(newnodes)
		return newnodes
}

func addNode(node string, nodes []string) []string {
	return append(nodes, node)
}

func whichNodeOwnsKey(key string, orderedNodes []string) string {
	var maxValue uint64 = 0
	var max string
	for i := 0; i < len(orderedNodes); i++ {
		var nodeHash = toHash(key, orderedNodes[i])
		if nodeHash > maxValue {
			max = orderedNodes[i]
			maxValue = nodeHash
		}
	}
	return max
}

func main() {
	fmt.Println(toHash("Soren Mathiasen"))	

	hasher := murmur3.New64()
	hasher.Write([]byte("Soren"))
	var hash = hasher.Sum64()
	fmt.Println(hash)


	
	hasher = murmur3.New64()
	hasher.Write([]byte("Joe"))
	hasher.Write([]byte("Soren"))
	hash = hasher.Sum64()
	fmt.Println("Single Joe and Soren", hash)
	fmt.Println("Single via function Joe and Soren", toHash("Joe", "Soren"))


	var nodes []string
	nodes = addNode("node1", nodes)
	nodes = addNode("node2", nodes)
	var myNode = whichNodeOwnsKey("key", nodes)

	nodes = removeNode(myNode, nodes)
	nodes = addNode(myNode, nodes)
	fmt.Println(myNode, whichNodeOwnsKey("key", nodes))
	//var orderedNodes []string

}
