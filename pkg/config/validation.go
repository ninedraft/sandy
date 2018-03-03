package sandyconfig

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"strings"
)

type parsedData struct {
	Tree   *toml.Tree
	Config Config
}

func (data *parsedData) Validate() (*Config, error) {
	if err := data.validateName(); err != nil {
		return nil, err
	}
	return &data.Config, nil
}

func (data *parsedData) validateName() error {
	if strings.TrimSpace(data.Config.Name) == "" {
		if data.Tree.Has("Name") {
			line := data.Tree.GetPosition("Name").Line
			return fmt.Errorf(`"Name" defined as zero string at line %d`, line)
		}
	}
	return nil
}
