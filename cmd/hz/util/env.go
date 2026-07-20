package util

import (
	"regexp"
)

func GetGOPATH() (gopath string, err error) { _ = "STUB: not implemented"; return "", nil }

func GetBuildGoPaths() []string { _ = "STUB: not implemented"; return nil }

var goModReg = regexp.MustCompile(`^\s*module\s+(\S+)\s*`)

func SearchGoMod(cwd string, recurse bool) (moduleName, path string, found bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func InitGoMod(module string) error { _ = "STUB: not implemented"; return nil }

func IsWindows() bool { _ = "STUB: not implemented"; return false }
