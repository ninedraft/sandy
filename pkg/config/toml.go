package sandyconfig

import (
	"bytes"
	"github.com/pelletier/go-toml"
)

func Parse(data []byte) (*parsedData, error) {
	decoder := toml.NewDecoder(bytes.NewReader(data))
	parsed := parsedData{}
	err := decoder.Decode(&parsed.Config)
	if err != nil {
		return nil, err
	}
	tree, err := toml.LoadBytes(data)
	if err != nil {
		return nil, err
	}
	parsed.Tree = tree
	return &parsed, nil
}
