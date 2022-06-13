package game

import (
	"io/ioutil"

	"github.com/pelletier/go-toml/v2"
)

type Configuration struct {
	Version  string
	Database DatabaseConfiguration
	FPS      int32
	Window   WindowConfiguration
}

type DatabaseConfiguration struct {
	DSN string
}

type WindowConfiguration struct {
	Width      int32
	Height     int32
	Title      string
	Fullscreen bool
}

var config *Configuration

func Config() *Configuration {
	if config != nil {
		return config
	}

	contents, err := ioutil.ReadFile("config/game.toml")
	err = toml.Unmarshal(contents, &config)

	if err != nil {
		panic(err)
	}

	return config
}
