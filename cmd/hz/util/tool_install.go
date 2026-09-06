package util

const ThriftgoMiniVersion = "v0.2.0"

func QueryVersion(exe string) (version string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ShouldUpdate(current, latest string) bool { _ = "STUB: not implemented"; return false }

func InstallAndCheckThriftgo() error { _ = "STUB: not implemented"; return nil }

func CheckCompiler(tool string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func CheckAndUpdateThriftgo() error { _ = "STUB: not implemented"; return nil }
