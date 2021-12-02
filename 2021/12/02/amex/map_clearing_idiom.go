package main

func clearNonFast(m map[string]int) {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// Do something with keys.
	// _ = keys
	for _, key := range keys {
		delete(m, key)
	}
}

func clearFast(m map[string]int) {
	for key := range m {
		delete(m, key)
	}
}
