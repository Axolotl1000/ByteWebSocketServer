package main

import "crypto/ed25519"

func Sign(privateKey []byte, message []byte) ([]byte, error) {
	sig := ed25519.Sign(privateKey, message)
	return sig, nil
}

func Verify(publicKey []byte, message []byte, signature []byte) bool {
	return ed25519.Verify(publicKey, message, signature)
}
