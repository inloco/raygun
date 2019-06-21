package raygun

import (
	"strings"
)

func GetCurrentStack() StackTrace {
	s := StackTrace{}
	Current(&s)

	offset := 1 // package github.com/kaeuferportal/stack2struct
	for strings.HasPrefix(s[offset].PackageName, PackageName) {
		offset++
	}

	return s[offset:]
}
