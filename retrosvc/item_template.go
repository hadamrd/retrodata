package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) ItemTemplates(ctx context.Context) (map[int]retro.ItemTemplate, error) {
	return svc.storer.ItemTemplates(ctx)
}
