func TestCode(t *testing.T) {
	tests := []struct {
		name string
		want any
		in   any
	}{{ /* Add your test cases here */ }, { /* More tests */ }}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := code(tt.in)
			if tt.wantErr != "" {
				return
			}
			if got != tt.want {
				t.Fatalf("Mismatch:\n\tGot:  %#v\n\tWant: %#v", got, tt.want)
			}
		})
	}
}
