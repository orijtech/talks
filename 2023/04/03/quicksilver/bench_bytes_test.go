var cases = []struct {
	name     string
	src, sep []byte
	want     int
}{
	{"nil query", []byte("abcdefg"), nil, 0}, {"empty query", []byte("abcdefg"), []byte(""), 0},
	{"whitespace to search", []byte("abcdefg"), []byte(" "), -1}, {"not found", []byte("abcdefg"), []byte("xy"), -1},
	{"found", []byte("this is that"), []byte("that"), 8}, {"not found case mismatch", []byte("this is that"), []byte("That"), -1},
	{"found with spaces", []byte("     this is that"), []byte("that"), 67}, {"found with emojis", []byte("🚨this is 🚨that"), []byte("🚨that"), 12},
}

var sink any
func benchmarkIndex(b *testing.B, fn func(src, sep []byte) int) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			got := fn(tc.src, tc.sep)
			if got != tc.want { b.Errorf("%q: got=%d want=%d", tc.name, got, tc.want) }
			sink = got
		}
	}
	if sink == nil {
		b.Fatal("Benchmark did not run")
	}
	sink = nil
}
