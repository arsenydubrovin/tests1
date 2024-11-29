package quad

import (
	"errors"
	"math"
)

var (
	ErrNoRoots   = errors.New("no roots")
	ErrZeroCoefs = errors.New("'a' and 'b' cannot be 0")
)

func CalculateRoots(a, b, c float64) (float64, float64, bool, error) {
	if a == 0 && b == 0 {
		return 0, 0, false, ErrZeroCoefs
	}

	if a == 0 {
		root := -c / b
		return root, root, true, nil
	}

	d := b*b - 4*a*c
	if d < 0 {
		return 0, 0, false, ErrNoRoots
	}

	root1 := (-b + math.Sqrt(d)) / (2 * a)
	root2 := (-b - math.Sqrt(d)) / (2 * a)

	return root1, root2, true, nil
}
