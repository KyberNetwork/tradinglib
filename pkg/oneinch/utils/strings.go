package utils

import "strings"

func Trim0x(s string) string {
	return strings.TrimPrefix(s, "0x")
}

func Add0x(s string) string {
	return "0x" + s
}

// IsHexString checks if a string is a valid hex string.
// valid hex string must start with 0x and contain only hex characters in lowercase.
func IsHexString(s string) bool {
	if len(s) < 2 || s[:2] != "0x" {
		return false
	}
	for _, c := range s[2:] {
		if !isHexChar(c) {
			return false
		}
	}
	return true
}

// isHexChar checks if a character is a valid hex character.
func isHexChar(c rune) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f')
}
