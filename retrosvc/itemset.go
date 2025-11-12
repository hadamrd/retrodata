package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) ItemSets(ctx context.Context) (map[int]retro.ItemSet, error) {
	return svc.storer.ItemSets(ctx)
}
