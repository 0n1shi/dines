package disasm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDines(t *testing.T) {
	assert := assert.New(t)

	type Expected struct {
		dines *Dines
		err   error
	}

	var tests = []struct {
		input    string
		expected Expected
	}{
		{
			input: "example.rom",
			expected: Expected{
				dines: &Dines{romFilePath: "example.rom"},
				err:   nil,
			},
		},
		{
			input: "",
			expected: Expected{
				dines: nil,
				err:   errors.New("ROM file path is empty"),
			},
		},
	}

	for _, test := range tests {
		d, e := NewDines(test.input)
		assert.Equal(d, test.expected.dines)
		assert.Equal(e, test.expected.err)
	}
}
