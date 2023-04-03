func bytesIndex(src, query []byte) (idx int) {
	if len(query) == 0 {
		return 0
	}

	for i := 0; i < len(src); i++ {
		if len(src[i:]) < len(query) {
			break
		}

		if bytes.Equal(src[i:i+len(query)], query) {
			return i
		}
	}
	return -1
}
