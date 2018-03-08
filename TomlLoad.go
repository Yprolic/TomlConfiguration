package TomlConfiguration

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io"
	"os"
	"path/filepath"
)

var (
	ErrSourceNotSet = errors.New("config path or reader is not set")
	ErrFileNotFound = errors.New("config file not found")
)

type TOMLLoader struct {
	Path   string
	Reader io.Reader
}

func (t *TOMLLoader) Load(s interface{}) error {
	var r io.Reader

	if t.Reader != nil {
		r = t.Reader
	} else if t.Path != "" {
		file, err := getConfig(t.Path)
		if err != nil {
			return err
		}
		defer file.Close()
		r = file
	} else {
		return ErrSourceNotSet
	}

	if _, err := toml.DecodeReader(r, s); err != nil {
		return err
	}

	return nil
}

func getConfig(path string) (*os.File, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := path
	if !filepath.IsAbs(path) {
		configPath = filepath.Join(pwd, path)
	}

	// check if file with combined path is exists(relative path)
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		return os.Open(configPath)
	}

	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, ErrFileNotFound
	}
	return f, err
}
