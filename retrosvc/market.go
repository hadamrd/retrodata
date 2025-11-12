package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) Markets(ctx context.Context) (map[string]retro.Market, error) {
	return svc.storer.Markets(ctx, svc.gameServerId)
}
