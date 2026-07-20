package config

import (
	"os/exec"
)

func lookupTool(idlType string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func link(src, dst string) error { _ = "STUB: not implemented"; return nil }

func BuildPluginCmd(args *Argument) (*exec.Cmd, error) { _ = "STUB: not implemented"; return nil, nil }

func (arg *Argument) GetThriftgoOptions() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
