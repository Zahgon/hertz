package meta

import (
	"path/filepath"
	"runtime"
)

const Version = "v0.9.7"

const DefaultServiceName = "hertz_service"

type Mode int

const SysType = runtime.GOOS

const WindowsOS = "windows"

const EnvPluginMode = "HERTZ_PLUGIN_MODE"

const (
	CmdUpdate = "update"
	CmdNew    = "new"
	CmdModel  = "model"
	CmdClient = "client"
)

const (
	IdlThrift = "thrift"
	IdlProto  = "proto"
)

const (
	TpCompilerThrift = "thriftgo"
	TpCompilerProto  = "protoc"
)

const (
	ProtocPluginName = "protoc-gen-hertz"
	ThriftPluginName = "thrift-gen-hertz"
)

const (
	LoadError           = 1
	GenerateLayoutError = 2
	PersistError        = 3
	PluginError         = 4
)

const (
	ModelDir   = "biz" + string(filepath.Separator) + "model"
	RouterDir  = "biz" + string(filepath.Separator) + "router"
	HandlerDir = "biz" + string(filepath.Separator) + "handler"
)

type Backend string

const (
	BackendGolang Backend = "golang"
)

const (
	SetBodyParam = "setBodyParam(req).\n"
)

const TheUseOptionMessage = "'model code' is not generated due to the '-use' option"

const AddThriftReplace = "do not generate 'go.mod', please add 'replace github.com/apache/thrift => github.com/apache/thrift v0.13.0' to your 'go.mod'"
