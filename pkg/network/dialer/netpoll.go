//go:build (amd64 || arm64) && (linux || darwin)

package dialer

import (
	"os"
	"strconv"

	"github.com/cloudwego/hertz/pkg/network/netpoll"
)

func init() {
	if v, _ := strconv.ParseBool(os.Getenv("HERTZ_NO_NETPOLL")); !v {
		defaultDialer = netpoll.NewDialer()
	}
}
