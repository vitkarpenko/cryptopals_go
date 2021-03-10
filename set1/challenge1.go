package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBytes(input string) (result []byte) {
	bytes := []byte(input)
	result = make([]byte, hex.DecodedLen(len(bytes)))
	_, err := hex.Decode(result, bytes)
	if err != nil {
		log.Fatalln("Incorrect hex string to decode.")
	}
	return
}

func Base64ToString(input []byte) (result string) {
	result = base64.StdEncoding.EncodeToString(input)
	return
}
