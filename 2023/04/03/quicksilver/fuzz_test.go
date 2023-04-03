func FuzzIt(f *testing.F) {
	// 1. Seed the fuzzer by adding in known/initial inputs.
	f.Add("abcd")
	f.Add("12adbe")
	f.Add("      ")

	// 2. Now run the fuzzer.
	f.Fuzz(func(t *testing.T, input string) {
		if err := myCode(input); err != nil { /* Handle the error appropriately */ } })
}
