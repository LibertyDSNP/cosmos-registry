package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DelegationKeyPrefix is the prefix to retrieve all Delegations
	DelegationKeyPrefix = "Delegations/value/"
	DelegationIndex     = "Delegation/index/"
)

// DelegationKey returns the store key to retrieve a Delegations from the index fields
func DelegationKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
