package set1

import (
	"sort"
)

var MostFrequentEnglishLetters = map[rune]struct{}{
	'e': {}, 'a': {}, 'r': {}, 'i': {}, 'o': {},
	't': {}, 'n': {}, 's': {}, 'l': {}, 'c': {},
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type RuneIntPair struct {
	Rune rune
	Int  int
}

type CipherRaitingPair struct {
	Cipher  int
	Raiting int
}

func findMostFrequentLetters(message string) map[rune]struct{} {
	counts := make(map[rune]int)
	for _, letter := range message {
		counts[letter] += 1
	}

	var countsSlice []RuneIntPair
	for k, v := range counts {
		countsSlice = append(countsSlice, RuneIntPair{Rune: k, Int: v})
	}

	sort.Slice(countsSlice, func(i, j int) bool {
		return countsSlice[i].Int > countsSlice[j].Int
	})

	result := make(map[rune]struct{})
	for _, kv := range countsSlice[:min(len(countsSlice), 10)] {
		result[kv.Rune] = struct{}{}
	}

	return result
}

func scoreStringAsEnglish(message string) (score int) {
	frequentLetters := findMostFrequentLetters(message)
	for letter := range frequentLetters {
		if _, ok := MostFrequentEnglishLetters[letter]; ok {
			score += 1
		}
	}
	return
}

func XORMessageByCipher(messageBytes []byte, cipher byte) (result []byte) {
	for i := range messageBytes {
		result = append(result, messageBytes[i]^cipher)
	}
	return
}

func DecryptXORedMessage(message string) map[int]string {
	messageBytes := HexToBytes(message)

	cipherRaitings := make(map[int]int)
	for cipher := range [256]int{} {
		xoredMessage := XORMessageByCipher(messageBytes, byte(cipher))
		cipherRaitings[cipher] = scoreStringAsEnglish(string(xoredMessage))
	}

	var raitingsSlice []CipherRaitingPair
	for k, v := range cipherRaitings {
		raitingsSlice = append(raitingsSlice, CipherRaitingPair{Cipher: k, Raiting: v})
	}

	sort.Slice(raitingsSlice, func(i, j int) bool {
		return raitingsSlice[i].Raiting > raitingsSlice[j].Raiting
	})

	result := make(map[int]string, 10)
	for _, cipherResult := range raitingsSlice[:10] {
		result[cipherResult.Cipher] = string(XORMessageByCipher(messageBytes, byte(cipherResult.Cipher)))
	}
	return result
}
