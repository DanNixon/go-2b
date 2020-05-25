package powerbox

import "github.com/DanNixon/go-2b/pkg/types"

type Powerbox interface {
	Reset() (types.Status, error)
	Kill() (types.Status, error)
	Set(s types.Settings) (types.Status, error)
	Get() (types.Status, error)
}
