package suite

import "github.com/cloudwego/hertz/pkg/protocol/client"

type ClientFactory interface {
	NewHostClient() (hc client.HostClient, err error)
}
