package set1

import (
	"sort"
)

var MostFrequentEnglishLetters = map[rune]int{
	'e': 10, 'a': 9, 'r': 8, 'i': 7, 'o': 6,
	't': 5, 'n': 4, 's': 3, 'l': 2, 'c': 1,
}

type CipherRaitingPair struct {
	Cipher  int
	Raiting int
}

func scoreStringAsEnglish(message string) (score int) {
	for _, letter := range message {
		if letterScore, ok := MostFrequentEnglishLetters[letter]; ok {
			score += letterScore
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

func DecryptXORedMessage(message string) string {
	messageBytes := HexToBytes(message)

	cipherRaitings := make(map[int]int)
	for cipher := range [256]int{} {
		xoredMessage := XORMessageByCipher(messageBytes, byte(cipher))
		cipherRaiting := scoreStringAsEnglish(string(xoredMessage))
		// Ignore the weaklings.
		if cipherRaiting >= 80 {
			cipherRaitings[cipher] = cipherRaiting
		}
	}

	var raitingsSlice []CipherRaitingPair
	for k, v := range cipherRaitings {
		raitingsSlice = append(raitingsSlice, CipherRaitingPair{Cipher: k, Raiting: v})
	}

	sort.Slice(raitingsSlice, func(i, j int) bool {
		return raitingsSlice[i].Raiting > raitingsSlice[j].Raiting
	})

	if len(raitingsSlice) != 0 {
		return string(XORMessageByCipher(messageBytes, byte(raitingsSlice[0].Cipher)))
	} else {
		return ""
	}
}
