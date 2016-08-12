// +build !cgo

package tcp_diag

import (
	"time"
	"errors"

	"github.com/elastic/beats/packetbeat/procs"
)

func GetSocketInfo(pollInterval time.Duration, socketInfo chan<- *procs.SocketInfo) error {
	return errors.New("only implemented on linux")
}
