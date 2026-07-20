//go:build !(amd64 || arm64) || stdjson

package json

import "encoding/json"

const Name = "encoding/json"

var (
	Marshal = json.Marshal

	Unmarshal = json.Unmarshal

	MarshalIndent = json.MarshalIndent

	NewDecoder = json.NewDecoder

	NewEncoder = json.NewEncoder
)
