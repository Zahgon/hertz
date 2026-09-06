package util

import (
	"go/ast"
	"go/token"
)

func AddImport(file, alias, impt string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddImportForContent(fileContent []byte, alias, impt string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addImport(fset *token.FileSet, f *ast.File, alias, impt string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
