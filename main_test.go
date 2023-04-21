package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type encryptDecryptData struct {
	key        string
	iv         string
	plaintext  []byte
	ciphertext []byte
}

func TestEncryptDecryptWithValidData(t *testing.T) {
	tc := encryptDecryptData{
		key:       "38695e9bc3b84e6984ad804e69fa07489e427734802c1c91e2e326c1222cef44",
		iv:        "bcb5dd1dfa03632bf408509cfd7ad1d9",
		plaintext: []byte("secret_pass"),
	}

	key, iv, _ := encodeKeyAndIv(tc.key, tc.iv)
	ciphertext, _ := encrypt(tc.plaintext, key, iv)
	plaintext, _ := decrypt(ciphertext, key, iv)

	assert.Equal(t, plaintext, tc.plaintext)
}

func TestEncryptDecryptWithInvalidKey(t *testing.T) {
	tc := encryptDecryptData{
		key:       "invalid_key",
		iv:        "bcb5dd1dfa03632bf408509cfd7ad1d9",
		plaintext: []byte("secret_pass"),
	}

	_, _, err := encodeKeyAndIv(tc.key, tc.iv)
	assert.NotNil(t, err)
}

func TestEncryptDecryptWithInvalidIv(t *testing.T) {
	tc := encryptDecryptData{
		key:       "38695e9bc3b84e6984ad804e69fa07489e427734802c1c91e2e326c1222cef44",
		iv:        "invalid_iv",
		plaintext: []byte("secret_pass"),
	}

	_, _, err := encodeKeyAndIv(tc.key, tc.iv)
	assert.NotNil(t, err)
}

func TestEncryptDecryptWithValidCipherText(t *testing.T) {
	tc := encryptDecryptData{
		key:        "38695e9bc3b84e6984ad804e69fa07489e427734802c1c91e2e326c1222cef44",
		iv:         "bcb5dd1dfa03632bf408509cfd7ad1d9",
		ciphertext: []byte("invalid_cipher_text"),
	}

	key, iv, _ := encodeKeyAndIv(tc.key, tc.iv)

	_, err := decrypt(tc.ciphertext, key, iv)
	assert.NotNil(t, err)
}
