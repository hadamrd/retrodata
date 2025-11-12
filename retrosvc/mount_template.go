package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) MountTemplates(ctx context.Context) (map[int]retro.MountTemplate, error) {
	return svc.storer.MountTemplates(ctx)
}
