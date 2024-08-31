package main

import (
	"golang.org/x/term"
)

var maxWidth int

func init() {
	columns, _, err := term.GetSize(1)
	if err != nil {
		maxWidth = 72
	} else {
		maxWidth = columns - 8
	}
}

func trimToEllipsis(s string) string {
	if len(s) > maxWidth {
		return s[:maxWidth] + "..."
	}
	return s
}
