package data

import (
	"encoding/hex"
)

// hexBytes is a byte slice that is unmarshalled from a hexstring.
type hexBytes []byte

func (b *hexBytes) UnmarshalText(text []byte) error {
	*b = make([]byte, hex.DecodedLen(len(text)))
	_, err := hex.Decode(*b, text)
	return err
}

type solcVersion struct {
	Path   string
	Sha256 [32]byte
}
