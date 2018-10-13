func fib(n int) int {
	if n < 2 {
		return n
	}

	ch := make(chan int)
	go fibGo(n-1, ch)
	go fibGo(n-2, ch)

	x := <-ch
	y := <-ch
	return x + y
}

func fibGo(n int, ch chan int) {
	ch <- fib(n)
}
