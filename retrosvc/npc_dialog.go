package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) NPCDialogs(ctx context.Context) (map[int]retro.NPCDialog, error) {
	return svc.storer.NPCDialogs(ctx)
}
