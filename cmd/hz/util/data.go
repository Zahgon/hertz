package util

import (
	"regexp"
)

func CopyStringSlice(from, to *[]string) { _ = "STUB: not implemented"; return }

func CopyString2StringMap(from, to map[string]string) { _ = "STUB: not implemented"; return }

func PackArgs(c interface{}) (res []string, err error) { _ = "STUB: not implemented"; return nil, nil }

func UnpackArgs(args []string, c interface{}) error { _ = "STUB: not implemented"; return nil }

func MapForm(input []string) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetFirstKV(m map[string][]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func ToCamelCase(name string) string { _ = "STUB: not implemented"; return "" }

func ToSnakeCase(name string) string { _ = "STUB: not implemented"; return "" }

func unifyPath(path string) string { _ = "STUB: not implemented"; return "" }

func BaseName(include, subFixToTrim string) string { _ = "STUB: not implemented"; return "" }

func BaseNameAndTrim(include string) string { _ = "STUB: not implemented"; return "" }

func SplitPackageName(pkg, subFixToTrim string) string { _ = "STUB: not implemented"; return "" }

func SplitPackage(pkg, subFixToTrim string) string { _ = "STUB: not implemented"; return "" }

func ImportToSanitizedPath(path string) string { _ = "STUB: not implemented"; return "" }

func ToVarName(paths []string) string { _ = "STUB: not implemented"; return "" }

func SplitGoTags(input string) []string { _ = "STUB: not implemented"; return nil }

func SubPackage(mod, dir string) string { _ = "STUB: not implemented"; return "" }

func SubDir(root, subPkg string) string { _ = "STUB: not implemented"; return "" }

var (
	uniquePackageName        = map[string]bool{}
	uniqueMiddlewareName     = map[string]bool{}
	uniqueHandlerPackageName = map[string]bool{}
)

func GetPackageUniqueName(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetMiddlewareUniqueName(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetHandlerPackageUniqueName(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getUniqueName(name string, uniqueNameSet map[string]bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var validFuncReg = regexp.MustCompile("[_0-9a-zA-Z]")

func ToGoFuncName(s string) string { _ = "STUB: not implemented"; return "" }
