package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cases := []struct {
		setup     func() (string, error)
		base      func() (string, error)
		assertion func(string)
		cleanup   func(string) error
	}{
		// Default.
		{
			setup: func() (string, error) {
				memory := os.Getenv("ENV")
				configCache = nil
				envCache = ""
				err := os.Setenv("ENV", "")
				if err != nil {
					return "", err
				}
				return memory, nil
			},
			base: func() (string, error) {
				cfg, err := Config()
				if err != nil {
					return "", err
				}
				return cfg["envname"], nil
			},
			assertion: func(value string) {
				assert.Equal(t, "default", value)
			},
			cleanup: func(memory string) error {
				return os.Setenv("ENV", memory)
			},
		},
		// Test.
		{
			setup: func() (string, error) {
				memory := os.Getenv("ENV")
				configCache = nil
				envCache = ""
				err := os.Setenv("ENV", "test")
				if err != nil {
					return "", err
				}
				return memory, nil
			},
			base: func() (string, error) {
				cfg, err := Config()
				if err != nil {
					return "", err
				}
				return cfg["envname"], nil
			},
			assertion: func(value string) {
				assert.Equal(t, "test", value)
			},
			cleanup: func(memory string) error {
				return os.Setenv("ENV", memory)
			},
		},
		// Production.
		{
			setup: func() (string, error) {
				memory := os.Getenv("ENV")
				configCache = nil
				envCache = ""
				err := os.Setenv("ENV", "production")
				if err != nil {
					return "", err
				}
				return memory, nil
			},
			base: func() (string, error) {
				cfg, err := Config()
				if err != nil {
					return "", err
				}
				return cfg["envname"], nil
			},
			assertion: func(value string) {
				assert.Equal(t, "production", value)
			},
			cleanup: func(memory string) error {
				return os.Setenv("ENV", memory)
			},
		},
	}

	for _, c := range cases {
		memory, err := c.setup()
		assert.Nil(t, err)
		value, err := c.base()
		assert.Nil(t, err)
		c.assertion(value)
		err = c.cleanup(memory)
		assert.Nil(t, err)
	}
}
