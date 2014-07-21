package hashing

import (
	"github.com/spaolacci/murmur3"
)


func toHash(value ...string) uint64 {
	hasher := murmur3.New64()
	for _, v := range value {
    	hasher.Write([]byte(v))
    }	
	var hash = hasher.Sum64()
	return hash
}
