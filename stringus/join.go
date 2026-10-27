package main

import "strings"

func Join(parts []string, sep string) string {
	var sb strings.Builder
	for i, part := range parts {
		sb.WriteString(part)
		if i < len(parts)-1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}