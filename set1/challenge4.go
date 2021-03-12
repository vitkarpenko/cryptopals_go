package set1

import (
	"bufio"
	"log"
	"os"
)

func loadStringsFromFile(path string) (strings []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %v!\n", path)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	if file.Close() != nil {
		log.Fatalf("Can't close file %v!\n", path)
	}
	return
}

func FindEncryptedMessage() (decrypted string) {
	messages := loadStringsFromFile("set1/data/challenge4.txt")
	for _, message := range messages {
		if decrypted := DecryptXORedMessage(message); len(decrypted) != 0 {

		}
	}
	return
}
