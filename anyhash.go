package anyhash

import (
	"encoding/binary"
	"fmt"

	xxhash "github.com/OneOfOne/xxhash/native"
)

// Dump will hash your data
func Dump(data interface{}, _ int) (result []byte) {
	result = make([]byte, 8)
	b := []byte(fmt.Sprintf("%#v", data))
	res := xxhash.Checksum64(b)
	binary.LittleEndian.PutUint64(result, res)
	return result
}

// Hash hashes data to hashes
func Hash(data interface{}) []byte {
	return Dump(data, 0)
}
