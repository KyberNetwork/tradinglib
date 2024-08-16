package limitorder

import "strings"

func trim0x(s string) string {
	return strings.TrimPrefix(s, "0x")
}
