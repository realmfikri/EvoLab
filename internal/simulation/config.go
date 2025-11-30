package simulation

import (
	"errors"
	"time"
)

// Config describes basic simulation parameters. In a real application this
// would be loaded from disk or a remote configuration service.
type Config struct {
	Steps    int
	Timestep time.Duration
	Render   bool
}

// LoadConfig returns a Config populated with default values. The provided path
// is accepted for future extension but is not required for the initial module
// scaffolding.
func LoadConfig(path string) (Config, error) {
	if path == "" {
		return Config{}, errors.New("config path cannot be empty")
	}

	return Config{
		Steps:    10,
		Timestep: 16 * time.Millisecond,
		Render:   true,
	}, nil
}
