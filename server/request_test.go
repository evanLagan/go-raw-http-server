package server

import (
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	tests := [] struct {
		name	string
		input	string
		want	Request
		expectErr bool
	}{
		{
			name: "Valid GET request",
			input: "GET / HTTP/1.1\n",
			want: Request{
				Method: "GET",
				Path:	"/",
				Version: "HTTP/1.1",
			},
			expectErr: false,
		},
		{
			name:	"Missing parts",
			input:	"GET /\n",
			expectErr:	true,
		},
		{
			name:	"Empty string",
			input:	"",
			expectErr: true,
		},
		{
			name:	"Only spaces",
			input:	"     ",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			got, err := parseRequestLine(tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("expected error, got nil")
			} else if !tt.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.expectErr && got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}

}