package gallery

import "context"

type Gallery interface {
	Start() error
	Stop(ctx context.Context) error
}
