package flashblock

import (
	"bytes"
	"io"
	
	"github.com/andybalholm/brotli"
)

func DecompressBrotli(input []byte) ([]byte, error) {
	reader := brotli.NewReader(bytes.NewReader(input))
	return io.ReadAll(reader)
}
