package quad

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRoots(t *testing.T) {
	tests := []struct {
		name      string
		a, b, c   float64
		expected1 float64
		expected2 float64
		hasRoots  bool
		err       error
	}{
		{
			name:      "different roots",
			a:         1,
			b:         2,
			c:         0,
			expected1: 0,
			expected2: -2,
			hasRoots:  true,
			err:       nil,
		},
		{
			name:      "equal roots",
			a:         2,
			b:         -4,
			c:         2,
			expected1: 1,
			expected2: 1,
			hasRoots:  true,
			err:       nil,
		},
		{
			name:      "zero 'a'",
			a:         0,
			b:         2,
			c:         1,
			expected1: -0.5,
			expected2: -0.5,
			hasRoots:  true,
			err:       nil,
		},
		{
			name:     "zero 'a' and 'b'",
			a:        0,
			b:        0,
			c:        1,
			hasRoots: false,
			err:      ErrZeroCoefs,
		},
		{
			name:     "no roots",
			a:        1,
			b:        0,
			c:        1,
			hasRoots: false,
			err:      ErrNoRoots,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root1, root2, hasRoots, err := CalculateRoots(tt.a, tt.b, tt.c)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.hasRoots, hasRoots)
			if tt.hasRoots {
				assert.Equal(t, tt.expected1, root1)
				assert.Equal(t, tt.expected2, root2)
			}
		})
	}
}
