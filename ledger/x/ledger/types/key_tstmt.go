package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TstmtKeyPrefix is the prefix to retrieve all Tstmt
	TstmtKeyPrefix = "Tstmt/value/"
)

// TstmtKey returns the store key to retrieve a Tstmt from the index fields
func TstmtKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
