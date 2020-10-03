package powerbox

import (
	"log"
	"time"

	"github.com/DanNixon/go-2b/pkg/types"
)

type Powerbox interface {
	Reset() (types.Status, error)
	Kill() (types.Status, error)
	Set(s types.Settings) (types.Status, error)
	Get() (types.Status, error)
}

type PowerboxTimeout struct {
	Timeout  time.Duration
	Timer    *time.Timer
	KillFunc func() (types.Status, error)
}

func NewPowerboxTimeout(timeout time.Duration, f func() (types.Status, error)) *PowerboxTimeout {
	t := PowerboxTimeout{
		Timeout:  timeout,
		KillFunc: f,
	}
	t.Timer = time.AfterFunc(timeout, t.onTimeout)
	return &t
}

func (t *PowerboxTimeout) Ping() {
	t.Timer.Stop()
	t.Timer.Reset(t.Timeout)
}

func (t *PowerboxTimeout) onTimeout() {
	t.KillFunc()
	log.Printf("Timeout elapsed, powerbox output killed")
	t.Timer.Reset(t.Timeout)
}
