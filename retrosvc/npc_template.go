package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) NPCTemplates(ctx context.Context) (map[int]retro.NPCTemplate, error) {
	return svc.storer.NPCTemplates(ctx)
}
