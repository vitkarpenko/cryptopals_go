package set1

import (
	"bytes"
)

func XORWithRepeatingKey(messageBytes []byte, key []byte) (result []byte) {
	cycledCipher := bytes.Repeat(key, len(messageBytes)/len(key)+1)
	for i := range messageBytes {
		result = append(result, messageBytes[i]^cycledCipher[i])
	}
	return
}
