package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RegistrationKeyPrefix is the prefix to retrieve all Registration
	RegistrationKeyPrefix = "Registration/value/"
	DsnpIdKey             = "Dsnp/value/"
)

// RegistrationKey returns the store key to retrieve a Registration from the index fields
func RegistrationKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
