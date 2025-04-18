package export_service

import (
	"fmt"
	"github.com/eyepipe/eye/internal/pkg/container"

	"encoding/json"
)

func ExportScheme(c *container.Container) ([]byte, error) {
	return json.Marshal(c.Scheme)
}

func ImportScheme(c *container.Container, bytea []byte) error {
	var scheme container.Scheme
	err := json.Unmarshal(bytea, &scheme)
	if err != nil {
		return fmt.Errorf("failed to unmarshall scheme: %w", err)
	}

	c.Scheme = scheme
	return nil
}
