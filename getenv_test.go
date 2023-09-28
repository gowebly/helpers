package gowebly

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	// Set the needed env.
	err := os.Setenv("GOWEBLY_HELPERS_TEST_1", "test")
	if err != nil {
		return
	}

	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "env is found, not set fallback value",
			args: args{
				key:      "GOWEBLY_HELPERS_TEST_1",
				fallback: "fallback",
			},
			want: "test",
		},
		{
			name: "env is not found, set fallback value",
			args: args{
				key:      "GOWEBLY_HELPERS_TEST_2",
				fallback: "fallback",
			},
			want: "fallback",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Getenv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("Getenv() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Clearenv()
}
