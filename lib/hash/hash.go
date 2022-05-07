package hash

import "crypto/sha256"

func GetKeyHash(key string) string {
	hash := sha256.Sum256([]byte(key))
	return string(hash[:])
}
