package hex

import (
	"encoding/hex"
	"fmt"
)

// b returns a byte slice from a hexstring or panics if the hexstring does not
// represent a vaild byte slice.
func b(hexBytes string) []byte {
	if !has0xPrefix(hexBytes) {
		panic(fmt.Sprintf("hex bytes %q must have 0x prefix", hexBytes))
	}
	if len(hexBytes)%2 != 0 {
		panic(fmt.Sprintf("hex bytes %q must have even number of hex chars", hexBytes))
	}

	bytes, err := hex.DecodeString(hexBytes[2:])
	if err != nil {
		panic(err)
	}
	return bytes
}

// has0xPrefix validates hexStr begins with '0x' or '0X'.
func has0xPrefix(hexStr string) bool {
	return len(hexStr) >= 2 && hexStr[0] == '0' && hexStr[1] == 'x'
}
