package sb

import (
	"fmt"
	"strings"
)

func Concat(elements ...string) string {
	b := strings.Builder{}
	tz := 0
	for _, e := range elements {
		tz += len(e)
	}
	b.Grow(tz)
	for _, e := range elements {
		b.WriteString(e)
	}
	return b.String()
}

func ConcatAny(elements ...interface{}) string {
	tmp := make([]string, 0, len(elements))
	for _, v := range elements {
		tmp = append(tmp, fmt.Sprint(v))
	}
	return Concat(tmp...)
}
