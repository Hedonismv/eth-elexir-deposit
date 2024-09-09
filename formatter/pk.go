package formatter

import (
	"bufio"
	"os"
	"strings"
)

// PrivateKeyToHex This function converts default private key with 0x... prefix to secp256k1 private key
func PrivateKeyToHex(privateKey string) string {
	if strings.HasPrefix(privateKey, "0x") {
		privateKey = privateKey[2:]
		return privateKey

	}
	return privateKey
}

func ReadKeysFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var keys []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keys = append(keys, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return keys, nil
}
