package thrift

import (
	"os"

	"github.com/cloudwego/hertz/cmd/hz/config"
	"github.com/cloudwego/hertz/cmd/hz/generator"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/thriftgo/parser"
	thriftgo_plugin "github.com/cloudwego/thriftgo/plugin"
)

type Plugin struct {
	req    *thriftgo_plugin.Request
	args   *config.Argument
	logger *logs.StdLogger
	rmTags []string
}

var debugPlugin = os.Getenv("HERTZ_DEBUG_PLUGIN") != ""

func NewPlugin(args *config.Argument, req *thriftgo_plugin.Request) *Plugin {
	_ = "STUB: not implemented"
	return nil
}

func (plugin *Plugin) Run() int { _ = "STUB: not implemented"; return 0 }

func (plugin *Plugin) Handle(args *config.Argument) (*thriftgo_plugin.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (plugin *Plugin) setLogger() { _ = "STUB: not implemented"; return }

func (plugin *Plugin) recvWarningLogger() string { _ = "STUB: not implemented"; return "" }

func (plugin *Plugin) recvVerboseLogger() string { _ = "STUB: not implemented"; return "" }

func (plugin *Plugin) handleRequest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (plugin *Plugin) parseArgs() (*config.Argument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (plugin *Plugin) initNameStyle() error { _ = "STUB: not implemented"; return nil }

func (plugin *Plugin) getPackageInfo() (*generator.HttpPackage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (plugin *Plugin) response(res *thriftgo_plugin.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (plugin *Plugin) InsertTag() ([]*thriftgo_plugin.Generated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (plugin *Plugin) GetResponse(files []generator.File, outputDir string) (*thriftgo_plugin.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTagString(f *parser.Field, rmTags []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
