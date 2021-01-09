package disasm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidROM(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    []byte
		Expected bool
	}{
		{
			Input:    MagicNumber,
			Expected: true,
		},
		{
			Input:    []byte{},
			Expected: false,
		},
		{
			Input:    []byte{0xde, 0xad, 0xbe, 0xef},
			Expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(isValidROM(test.Input), test.Expected)
	}
}
