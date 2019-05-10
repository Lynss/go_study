package ly

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

//  发送value，结束后关闭channel
func Walk(t *Tree, ch chan int) {
	sendValue(t, ch)
	close(ch)
}

//  递归向channel传值
func sendValue(t *Tree, ch chan int) {
	if t != nil {
		sendValue(t.Left, ch)
		ch <- t.Value
		sendValue(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	result := false
	for i := range c1 {
		result = result && (i == <-c2)
	}
	return result
}
