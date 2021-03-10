package set1

import (
	"log"
)

func FixedXOR(first, second string) (result []byte) {
	if len(first) != len(second) {
		log.Fatalln("Arguments length should be equal in FixedXOR function.")
	}
	firstBytes := HexToBytes(first)
	secondBytes := HexToBytes(second)
	result = make([]byte, len(firstBytes))
	for i := range firstBytes {
		result[i] = firstBytes[i] ^ secondBytes[i]
	}
	return
}
