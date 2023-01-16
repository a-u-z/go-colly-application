package main

import "strings"

func WriteStringByBuilder(len int, s ...string) string {
	var builder strings.Builder
	builder.Grow(len)
	for _, v := range s {
		builder.WriteString(v)
	}
	return builder.String()
}

func refinePrice(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "[售價]", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")

	return strings.ReplaceAll(s, "元", "")
}
