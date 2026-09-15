package greeting

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Gopher", "Hello, Gopher!"},
		{"", "Hello, World!"},
	}
	for _, tt := range tests {
		if got := Hello(tt.name); got != tt.want {
			t.Errorf("Hello(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestGoodbye(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Gopher", "Goodbye, Gopher!"},
		{"", "Goodbye, World!"},
	}
	for _, tt := range tests {
		if got := Goodbye(tt.name); got != tt.want {
			t.Errorf("Goodbye(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}
