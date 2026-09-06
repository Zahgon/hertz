package protobuf

import (
	"os"

	"github.com/cloudwego/hertz/cmd/hz/config"
	"github.com/cloudwego/hertz/cmd/hz/generator"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Plugin struct {
	*protogen.Plugin
	Package      string
	Recursive    bool
	OutDir       string
	ModelDir     string
	UseDir       string
	IdlClientDir string
	RmTags       RemoveTags
	PkgMap       map[string]string
	logger       *logs.StdLogger
}

type RemoveTags []string

func (rm *RemoveTags) Exist(tag string) bool { _ = "STUB: not implemented"; return false }

var debugPlugin = os.Getenv("HERTZ_DEBUG_PLUGIN") != ""

func (plugin *Plugin) Run() int { _ = "STUB: not implemented"; return 0 }

func (plugin *Plugin) setLogger() { _ = "STUB: not implemented"; return }

func (plugin *Plugin) recvWarningLogger() string { _ = "STUB: not implemented"; return "" }

func (plugin *Plugin) recvVerboseLogger() string { _ = "STUB: not implemented"; return "" }

func (plugin *Plugin) parseArgs(param string) (*config.Argument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (plugin *Plugin) Response(resp *pluginpb.CodeGeneratorResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func (plugin *Plugin) Handle(req *pluginpb.CodeGeneratorRequest, args *config.Argument) error {
	_ = "STUB: not implemented"
	return nil
}

func (plugin *Plugin) fixGoPackage(req *pluginpb.CodeGeneratorRequest, pkgMap map[string]string, trimGoPackage string) {
	_ = "STUB: not implemented"
	return
}

func (plugin *Plugin) fixModelPathAndPackage(pkg string) (impt string) {
	_ = "STUB: not implemented"
	return ""
}

func (plugin *Plugin) GenerateFiles(pluginPb *protogen.Plugin) error {
	_ = "STUB: not implemented"
	return nil
}

func (plugin *Plugin) GenerateFile(gen *protogen.Plugin, f *protogen.File) error {
	_ = "STUB: not implemented"
	return nil
}

func generateFile(gen *protogen.Plugin, file *protogen.File, rmTags RemoveTags) (*protogen.GeneratedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genMessage(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo, rmTags RemoveTags) error {
	_ = "STUB: not implemented"
	return nil
}

func genMessageFields(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo, rmTags RemoveTags) error {
	_ = "STUB: not implemented"
	return nil
}

func genMessageField(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo, field *protogen.Field, sf *structFields, rmTags RemoveTags) error {
	_ = "STUB: not implemented"
	return nil
}

func (plugin *Plugin) getIdlInfo(ast *descriptorpb.FileDescriptorProto, deps map[string]*descriptorpb.FileDescriptorProto, args *config.Argument) (*generator.HttpPackage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (plugin *Plugin) genHttpPackage(ast *descriptorpb.FileDescriptorProto, deps map[string]*descriptorpb.FileDescriptorProto, args *config.Argument) ([]generator.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
