package yamlfuzz

import (
	"gopkg.in/yaml.v3"
)

func Fuzz(data []byte) int {
	var n yaml.Node
	if err := yaml.Unmarshal(data, &n); err != nil {
		return 0  // Return 0 for any unmarshaling error (common case).
	}
	return 1  // Return 1 if the input is valid YAML.
}
