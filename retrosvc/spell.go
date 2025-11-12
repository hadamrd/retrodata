package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) Spells(ctx context.Context) (map[int]retro.Spell, error) {
	return svc.storer.Spells(ctx)
}
