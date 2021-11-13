package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetHeight(t *testing.T) {
	cases := map[string]struct {
		hill Hill
		want float32
	}{
		"t1": {
			hill: Hill{
				Length: 78.5,
				Slope:  0.3,
			},
			want: 22.5568,
		},
	}

	for name, tc := range cases {
		fmt.Printf("Test: %s\n", name)
		actual := tc.hill.GetHeight()
		assert.InDelta(t, tc.want, actual, 0.0001)
	}
}
