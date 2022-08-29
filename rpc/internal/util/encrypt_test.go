package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		origin string
	}{
		{
			origin: "123456",
		},
		{
			origin: "123456789..",
		},
	}

	for _, v := range tests {
		// test encrypt
		encryptedData := BcryptEncrypt(v.origin)
		result := BcryptCheck(v.origin, encryptedData)
		assert.Equal(t, result, true)
	}
}
