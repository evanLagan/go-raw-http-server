package server

import (
	"fmt"
	"strings"
)

type Request struct {
	Method string
	Path string
	Version string
}


// Splits the line by spaces
// Trims newline characters
// Validates the structure (must have 3 parts)
func parseRequestLine(line string) (Request, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return Request{}, fmt.Errorf("invalid request line: %s", line)
	}

	return Request{
		Method: parts[0],
		Path: parts[1],
		Version: strings.TrimSpace(parts[2]),
	}, nil
}