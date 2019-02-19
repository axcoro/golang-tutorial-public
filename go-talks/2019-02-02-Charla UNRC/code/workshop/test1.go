
func TestExtractRequestID(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractRequestID(tt.args.tags); got != tt.want {
				t.Errorf("ExtractRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}