package main

func main() {
	m := map[string]int{"ten": 10}
	ten := []byte("ten")
	// Expensive way.
	key := string(ten)
	println(m[key])

	// Cheap way.
	println(m[string(ten)])
}
