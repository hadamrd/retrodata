package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) GameMaps(ctx context.Context) (map[int]retro.GameMap, error) {
	return svc.storer.GameMaps(ctx)
}
