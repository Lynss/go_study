package ly

import "fmt"

func operate(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		b = a + b
		a = b - a
		return b - a
	}
}

func OutFib() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
