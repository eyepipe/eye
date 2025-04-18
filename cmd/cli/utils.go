package main

import (
	"fmt"
	"github.com/eyepipe/eye/internal/pkg/container"

	"encoding/json"
)

func MustMarshalJSON(v any) []byte {
	got, err := json.Marshal(container.SchemeDefault)
	if err != nil {
		panic(fmt.Errorf("failed to marshal to JSON: %v", err))
	}

	return got
}
