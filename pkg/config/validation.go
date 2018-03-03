package sandyconfig

import (
	"github.com/ninedraft/sandy/pkg/sanderr"
	"github.com/pelletier/go-toml"
	"strings"
)

var (
	ErrInvalidField = sanderr.NewErr(`invalid field`)
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
			return ErrInvalidField.
				AddStrF(`"Name" defined as zero string at line %d`, line)
		}
	}
	return nil
}
