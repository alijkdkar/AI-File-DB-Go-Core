package utility

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func EncryptAES(key []byte, plaintext string) (string, error) {
	// create cipher
	c, err := aes.NewCipher(key)

	if err != nil {
		CheckError(err)
		return "", nil
	}
	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out), nil
}

func DecryptAES(key []byte, ct string) (string, error) {

	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	if err != nil {
		CheckError(err)
		return "", err
	}

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])

	return s, nil
}

func CheckError(er error) {
	fmt.Println(er)
}

func GetKey(key string) []byte {

	if key == "" {
		if key1, err := ioutil.ReadFile("Hashkey.key"); err != nil {
			key = string(key1)
		}
	}

	return []byte(key)
}
