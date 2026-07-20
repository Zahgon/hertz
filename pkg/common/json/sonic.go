//go:build (amd64 || arm64) && !stdjson

package json

import "github.com/bytedance/sonic"

const Name = "sonic"

var (
	json = sonic.ConfigStd

	Marshal = json.Marshal

	Unmarshal = json.Unmarshal

	MarshalIndent = json.MarshalIndent

	NewDecoder = json.NewDecoder

	NewEncoder = json.NewEncoder
)
