package configuration

import (
	"os"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port int
}

func New(tableName string) Configuration {
	configuration := Configuration{}

	env := os.Getenv("ENVIRONMENT")

	if env == nil {
		env = "PROD"
	}

	path := "loko"

	err := gonfig.GetConf("path/to/myjonfile.json", &configuration)
}
