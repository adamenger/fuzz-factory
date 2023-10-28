package jwtfuzz

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
  "strings"
)

// Load the keys from disk
var (
	jwtTestDefaultKey       *rsa.PublicKey
	jwtTestRSAPrivateKey    *rsa.PrivateKey
	jwtTestEC256PublicKey   *ecdsa.PublicKey
	jwtTestEC384PublicKey   *ecdsa.PublicKey
	jwtTestEC512PublicKey   *ecdsa.PublicKey
	jwtTestEC256PrivateKey  *ecdsa.PrivateKey
	jwtTestEC384PrivateKey  *ecdsa.PrivateKey
	jwtTestEC512PrivateKey  *ecdsa.PrivateKey
	jwtTestEd25519PublicKey ed25519.PublicKey
	jwtTestEd25519PrivateKey ed25519.PrivateKey
	hmacTestKey []byte
	paddedKey               crypto.PublicKey
)

func loadRSAKeyFromDisk(path string) *rsa.PublicKey {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		panic(err)
	}
	return key
}

func loadECKeyFromDisk(path string) *ecdsa.PublicKey {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseECPublicKeyFromPEM(keyBytes)
	if err != nil {
		panic(err)
	}
	return key
}

func loadEdPublicKeyFromDisk(path string) ed25519.PublicKey {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseEdPublicKeyFromPEM(keyBytes)
	if err != nil {
		panic(err)
	}
	return key.(ed25519.PublicKey)
}

func loadHMACKeyFromDisk(path string) []byte {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return keyBytes
}

func init() {
	jwtTestDefaultKey = loadRSAKeyFromDisk("keys/sample_key.pub")
	jwtTestEC256PublicKey = loadECKeyFromDisk("keys/ec256-public.pem")
	jwtTestEC384PublicKey = loadECKeyFromDisk("keys/ec384-public.pem")
	jwtTestEC512PublicKey = loadECKeyFromDisk("keys/ec512-public.pem")
	jwtTestEd25519PublicKey = loadEdPublicKeyFromDisk("keys/ed25519-public.pem")
	hmacTestKey = loadHMACKeyFromDisk("keys/hmacTestKey")
	paddedKey = loadECKeyFromDisk("keys/examplePaddedKey-public.pem")
}

func FuzzParse(data []byte) int {
    // Too small a data slice won't be meaningful; early exit in such a case
    if len(data) < 3 {
        return -1
    }

    p := &jwt.Parser{}
    tokenString := string(data[2:])

    // Fuzz Parser Configuration based on the first two bytes of the data
    p.SkipClaimsValidation = (data[0]&1) == 1
    p.UseJSONNumber = (data[0]&2) == 2

    claimTypes := []jwt.Claims{
        jwt.MapClaims{},
    }

    keyFuncs := []jwt.Keyfunc{
        func(t *jwt.Token) (interface{}, error) { return jwtTestDefaultKey, nil },
        func(t *jwt.Token) (interface{}, error) { return jwtTestEC256PublicKey, nil },
        func(t *jwt.Token) (interface{}, error) { return jwtTestEC384PublicKey, nil },
        func(t *jwt.Token) (interface{}, error) { return jwtTestEC512PublicKey, nil },
        func(t *jwt.Token) (interface{}, error) { return jwtTestEd25519PublicKey, nil },
        func(t *jwt.Token) (interface{}, error) { return hmacTestKey, nil },
        func(t *jwt.Token) (interface{}, error) { return paddedKey, nil },
    }

    for _, keyFunc := range keyFuncs {
        for _, claimType := range claimTypes {
            _, err := p.ParseWithClaims(tokenString, claimType, keyFunc)

            if err != nil {
                if strings.Contains(err.Error(), "key is of invalid type") {
                    continue
                }
                return 1
            }
        }
    }

    // If parsing was successful across different configurations, return 0 (no interesting path)
    return 0
}

