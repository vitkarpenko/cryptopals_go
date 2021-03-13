package set1

import "log"

func hammingDistanceTwoStrings(first, second string) (distance int) {
	if len(first) != len(second) {
		log.Fatalln("Arguments length should be equal to compute Hamming distance.")
	}
	firstBytes := []byte(first)
	secondBytes := []byte(second)
	for i := range firstBytes {
		distance += hammingDistanceTwoBytes(firstBytes[i], secondBytes[i])
	}
	return
}

func hammingDistanceTwoBytes(first, second byte) (distance int) {
	for diff := first ^ second; diff > 0; distance++ {
		// Setting the lowest-order 1 to 0.
		diff = diff & (diff - 1)
	}
	return
}
