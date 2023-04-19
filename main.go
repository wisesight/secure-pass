package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

const hexKey = "TEMPLATE_KEY_32"
const hexIv = "TEMPLATE_IV_16"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./ttb-secure-pass (-e|--encrypt) plaintext OR (-d|--decrypt) ciphertext")
		os.Exit(1)
	}

	key, err := hex.DecodeString(hexKey)
	if err != nil {
		panic(err)
	}

	iv, err := hex.DecodeString(hexIv)
	if err != nil {
		panic(err)
	}

	if os.Args[1] == "-e" || os.Args[1] == "--encrypt" {
		plaintext := []byte(os.Args[2])
		ciphertext, err := encrypt(plaintext, key, iv)
		if err != nil {
			panic(err)
		}
		fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
	} else if os.Args[1] == "-d" || os.Args[1] == "--decrypt" {
		ciphertext, err := base64.StdEncoding.DecodeString(os.Args[2])
		if err != nil {
			panic(err)
		}
		decrypted, err := decrypt(ciphertext, key, iv)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(decrypted))
	} else {
		fmt.Println("Invalid option:", os.Args[1])
		os.Exit(1)
	}

}

func encrypt(plaintext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedPlaintext := pad(plaintext)

	ciphertext := make([]byte, len(paddedPlaintext))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	return ciphertext, nil
}

func decrypt(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	decrypted := make([]byte, len(ciphertext))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext)

	return unpad(decrypted), nil
}

func pad(plaintext []byte) []byte {
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func unpad(padded []byte) []byte {
	padding := int(padded[len(padded)-1])
	return padded[:len(padded)-padding]
}
