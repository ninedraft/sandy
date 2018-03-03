package sandyconfig

import (
	"bytes"
	"github.com/ninedraft/sandy/pkg/sanderr"
	"github.com/pelletier/go-toml"
)

var (
	ErrInvalidToml = sanderr.NewErr("invalid TOML file")
)

func Parse(data []byte) (*parsedData, error) {
	decoder := toml.NewDecoder(bytes.NewReader(data))
	parsed := parsedData{}
	err := decoder.Decode(&parsed.Config)
	if err != nil {
		err = ErrInvalidToml.
			AddError(err)
		return nil, err
	}
	tree, err := toml.LoadBytes(data)
	if err != nil {
		err = ErrInvalidToml.
			AddError(err)
		return nil, err
	}
	parsed.Tree = tree
	return &parsed, nil
}
