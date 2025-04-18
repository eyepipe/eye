package crypto2

import (
	"strings"
)

type StringPart struct {
	parts []string
}

func NewStringPart(input string, sep string) *StringPart {
	return &StringPart{
		parts: strings.Split(input, sep),
	}
}

// Get safely returns item from slice by index ot default value
func (s *StringPart) Get(index int) string {
	switch {
	case len(s.parts)-1 >= index:
		return s.parts[index]
	default:
		return ""
	}
}
