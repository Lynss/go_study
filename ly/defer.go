package ly

func hello() string {
	println("hello")
	return "!"
}

func DeferStack() {
	for i := 0; i < 10; i++ {
		defer println(i)
	}
}

func GreetWorld() {
	defer println(hello())
	println("world")
}
