package app

import (
	"github.com/cloudwego/hertz/cmd/hz/config"
	"github.com/urfave/cli/v2"
)

var globalArgs = config.NewArgument()

func New(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

func Update(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

func Model(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

func Client(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

func PluginMode() { _ = "STUB: not implemented"; return }

func Init() *cli.App { _ = "STUB: not implemented"; return nil }

func setLogVerbose(verbose bool) { _ = "STUB: not implemented"; return }

func GenerateLayout(args *config.Argument) error { _ = "STUB: not implemented"; return nil }

func TriggerPlugin(args *config.Argument) error { _ = "STUB: not implemented"; return nil }
