package helpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasePath(t *testing.T) {
	cases := []struct {
		setup     func() (map[string]string, error)
		base      func() string
		assertion func(string)
		cleanup   func(map[string]string) error
	}{
		// PROJECT_PATH present.
		{
			setup: func() (map[string]string, error) {
				memory := make(map[string]string)
				for _, k := range []string{"PROJECT_PATH"} {
					memory[k] = os.Getenv(k)
				}
				basePathCache = ""
				err := os.Setenv("PROJECT_PATH", "test-project-path")
				if err != nil {
					return nil, err
				}
				return memory, nil
			},
			base: func() string {
				return BasePath()
			},
			assertion: func(value string) {
				assert.Equal(t, "test-project-path", value)
			},
			cleanup: func(memory map[string]string) error {
				for k, v := range memory {
					err := os.Setenv(k, v)
					if err != nil {
						return err
					}
				}
				return nil
			},
		},
		// PROJECT_PATH missing, GOPATH present.
		{
			setup: func() (map[string]string, error) {
				memory := make(map[string]string)
				for _, k := range []string{"PROJECT_PATH", "GOPATH"} {
					memory[k] = os.Getenv(k)
				}
				basePathCache = ""
				err := os.Setenv("PROJECT_PATH", "")
				if err != nil {
					return nil, err
				}
				err = os.Setenv("GOPATH", "test-gopath")
				if err != nil {
					return nil, err
				}
				return memory, nil
			},
			base: func() string {
				return BasePath()
			},
			assertion: func(value string) {
				assert.Equal(t, "test-gopath/src/github.com/Oleg-Smal-git/boosters-trial", value)
			},
			cleanup: func(memory map[string]string) error {
				for k, v := range memory {
					err := os.Setenv(k, v)
					if err != nil {
						return err
					}
				}
				return nil
			},
		},
		// Both missing.
		{
			setup: func() (map[string]string, error) {
				memory := make(map[string]string)
				for _, k := range []string{"PROJECT_PATH", "GOPATH"} {
					memory[k] = os.Getenv(k)
				}
				basePathCache = ""
				err := os.Setenv("PROJECT_PATH", "")
				if err != nil {
					return nil, err
				}
				err = os.Setenv("GOPATH", "")
				if err != nil {
					return nil, err
				}
				return memory, nil
			},
			base: func() string {
				return BasePath()
			},
			assertion: func(value string) {
				assert.Equal(t, ".", value)
			},
			cleanup: func(memory map[string]string) error {
				for k, v := range memory {
					err := os.Setenv(k, v)
					if err != nil {
						return err
					}
				}
				return nil
			},
		},
	}

	for _, c := range cases {
		memory, err := c.setup()
		assert.Nil(t, err)
		result := c.base()
		c.assertion(result)
		err = c.cleanup(memory)
		assert.Nil(t, err)
	}
}
