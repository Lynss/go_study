package ly

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		distance := math.Abs(z*z - x)
		if distance < 0.01 {
			break
		}
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SwitchLike(a string) {
	switch a {
	case "a":
		fmt.Printf("shazi:%s;", a)
		fallthrough
	default:
		fmt.Println("dashazi")
	}
}

func SwitchWithoutCondition() {
	a := 1
	switch {
	case a > 0:
		fmt.Println("positive")
	case a < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")

	}
}
