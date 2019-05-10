package ly

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: v%", float64(e))
}

func SqrtWithError(x float64) (float64, error) {
	var z float64
	var e error = nil
	if x < 0 {
		e = ErrNegativeSqrt(x)
	} else {
		z = 1.0
		for {
			distance := math.Abs(z*z - x)
			if distance < 0.01 {
				break
			}
			z -= (z*z - x) / (2 * z)
		}
	}
	return z, e
}
