package internal

import "testing"

func TestColor(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		bgcolor Color
		want    string
	}{
		{
			name:    "empty string",
			text:    "",
			bgcolor: "\033[42m",
			want:    "",
		},
		{
			name:    "black string",
			text:    " ",
			bgcolor: "\033[42m",
			want:    "",
		},
		{
			name:    "no color",
			text:    "text",
			bgcolor: "",
			want:    "text",
		},
		{
			name:    "invalid color",
			text:    "text",
			bgcolor: "\x1b[0;37m",
			want:    "text",
		},
		{
			name:    "valid color",
			text:    "text",
			bgcolor: "\033[42m",
			want:    "\033[42m\x1b[0;37mtext\033[0m",
		},
	}

	colorizer := NewColorizer()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := colorizer.Colorize(tc.text, tc.bgcolor)
			if got != tc.want {
				t.Errorf("%s: got %q, want %q", tc.name, got, tc.want)
			}
		})
	}
}

func BenchmarkColorize(b *testing.B) {
	colorizer := NewColorizer()
	for i := 0; i < b.N; i++ {
		colorizer.Colorize("text", "\033[42m")
	}
}
