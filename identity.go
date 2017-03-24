package ipfs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	libp2p "gx/ipfs/QmPGxZ1DP2w45WcogpW1h43BvseXbfke9N91qotpoQcUeS/go-libp2p-crypto"
)

func IdentityKeyFromSeed(seed []byte, bits int) ([]byte, error) {
	hmac := hmac.New(sha256.New, []byte("OpenBazaar seed"))
	hmac.Write(seed)
	reader := bytes.NewReader(hmac.Sum(nil))
	sk, _, err := libp2p.GenerateKeyPairWithReader(libp2p.Ed25519, bits, reader)
	if err != nil {
		return nil, err
	}
	encodedKey, err := sk.Bytes()
	if err != nil {
		return nil, err
	}
	return encodedKey, nil
}
