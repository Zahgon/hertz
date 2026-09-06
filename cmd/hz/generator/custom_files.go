package generator

import (
	"bytes"
)

type FilePathRenderInfo struct {
	MasterIDLName  string
	GenPackage     string
	HandlerDir     string
	ModelDir       string
	RouterDir      string
	ProjectDir     string
	GoModule       string
	ServiceName    string
	MethodName     string
	HandlerGenPath string
}

type IDLPackageRenderInfo struct {
	FilePathRenderInfo
	ServiceInfos *HttpPackage
}

type CustomizedFileForMethod struct {
	*HttpMethod
	FilePath       string
	FilePackage    string
	ServiceInfo    *Service
	IDLPackageInfo *IDLPackageRenderInfo
}

type CustomizedFileForService struct {
	*Service
	FilePath       string
	FilePackage    string
	IDLPackageInfo *IDLPackageRenderInfo
}

type CustomizedFileForIDL struct {
	*IDLPackageRenderInfo
	FilePath    string
	FilePackage string
}

func (pkgGen *HttpPackageGenerator) genCustomizedFile(pkg *HttpPackage) error {
	_ = "STUB: not implemented"
	return nil
}

func renderFilePath(tplInfo *Template, filePathRenderInfo FilePathRenderInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func renderInsertKey(tplInfo *Template, data interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func renderImportTpl(tplInfo *Template, data interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renderAppendContent(tplInfo *Template, renderInfo interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func appendUpdateFile(tplInfo *Template, renderInfo interface{}, fileContent []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getInsertImportContent(tplInfo *Template, renderInfo interface{}, fileContent []byte) ([][2]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pkgGen *HttpPackageGenerator) genLoopService(tplInfo *Template, filePathRenderInfo FilePathRenderInfo, service *Service, idlPackageRenderInfo *IDLPackageRenderInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgGen *HttpPackageGenerator) genLoopMethod(tplInfo *Template, filePathRenderInfo FilePathRenderInfo, method *HttpMethod, service *Service, idlPackageRenderInfo *IDLPackageRenderInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgGen *HttpPackageGenerator) genSingleCustomizedFile(tplInfo *Template, filePathRenderInfo FilePathRenderInfo, idlPackageRenderInfo IDLPackageRenderInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func writeBytes(buf *bytes.Buffer, bytes ...[]byte) error { _ = "STUB: not implemented"; return nil }
