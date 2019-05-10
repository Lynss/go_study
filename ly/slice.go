package ly

func ShowRange() {
	a := make([][]int, 3)
	for i := range a {
		a[i] = make([]int, 3)
	}
	for i := range a {
		for j := range a[i] {
			a[i][j] = int(3)
		}
	}
	for i, v := range a {
		for j, k := range v {
			println(i, j, k)
		}
	}
}
