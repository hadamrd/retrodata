package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) EffectTemplates(ctx context.Context) (map[int]retro.EffectTemplate, error) {
	return svc.storer.EffectTemplates(ctx)
}
