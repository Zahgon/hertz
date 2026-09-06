package consts

import "time"

const (
	MaxSmallFileSize = 2 * 4096

	FSHandlerCacheDuration = 10 * time.Second

	FSCompressedFileSuffix    = ".hertz.gz"
	FsMinCompressRatio        = 0.8
	FsMaxCompressibleFileSize = 8 * 1024 * 1024
)
