package input_resolver_service

import (
	"strings"
)

func IsHttpURL(input string) bool {
	input = strings.ToLower(input)
	return strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://")
}
