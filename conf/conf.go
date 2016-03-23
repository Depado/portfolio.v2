package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// C is the exported global configuration variable
var C Conf

// Conf is the struct containing the configuration of the server
type Conf struct {
	Port  int  `yaml:"port"`
	Debug bool `yaml:"debug"`
}

// Load loads the given fp (file path) to the C global configuration variable.
func Load(fp string) error {
	var err error
	var conf []byte
	if conf, err = ioutil.ReadFile(fp); err != nil {
		return err
	}
	return yaml.Unmarshal(conf, &C)
}
