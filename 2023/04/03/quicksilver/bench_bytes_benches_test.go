func BenchmarkIndexNaive(b *testing.B) {
	benchmarkIndex(b, bytesIndexNaive)
}
func BenchmarkIndexStdlib(b *testing.B) {
	benchmarkIndex(b, bytes.Index)
}
