package internal

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateTrackingCode() string {
	bytes := make([]byte, 5)
	_, _ = rand.Read(bytes)

	return "SP-" + hex.EncodeToString(bytes)
}
