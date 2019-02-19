func TestExtractRequestID(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"OK", args{[]string{ "request_id:123", "asdasd:asdsad"}}, "request_id:123",
		}, {
			"Ignore", args{[]string{ "asdasd:1", "asdasd:2", "asdasd:3", "asdasd:4"}}, "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractRequestID(tt.args.tags); got != tt.want {
				t.Errorf("ExtractRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}
