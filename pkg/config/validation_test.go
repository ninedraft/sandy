package sandyconfig

import (
	"testing"
)

func TestNameValidation(test *testing.T) {
	validName := []byte(`
		Name = "name"
		`)
	invalidName := []byte(`
				Name = ""
		`)
	parse := func(tomlData []byte) error {
		data, err := Parse(tomlData)
		if err != nil {
			test.Fatalf("error while parsing test TOML data: %v", err)
		}
		return data.validateName()
	}
	if err := parse(validName); err != nil {
		test.Fatalf("error while validating test TOML data with VALID Name: %v", err)
	}
	if err := parse(invalidName); err == nil {
		test.Fatalf("nil error while validating TOML data with INVALID Name")
	} else {
		test.Log(err)
	}
}
