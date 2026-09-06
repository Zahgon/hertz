package utils

const (
	UNKNOWN_IP_ADDR = "-"
)

var localIP string

func LocalIP() string { _ = "STUB: not implemented"; return "" }

func getLocalIp() string { _ = "STUB: not implemented"; return "" }

func init() {
	localIP = getLocalIp()
}

func TLSRecordHeaderLooksLikeHTTP(hdr [5]byte) bool { _ = "STUB: not implemented"; return false }
