package ly

func ForCommon() {
	for i := 0; i < 10; i++ {
		println("wtf")
	}
}

func ForWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	println(sum)
}

func ForLoop() {
	sum := 1
	for {
		if sum >= 1000 {
			break
		}
		sum += sum
	}
	println(sum)
}
