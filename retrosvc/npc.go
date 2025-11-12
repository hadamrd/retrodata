package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) NPCs(ctx context.Context) (map[string]retro.NPC, error) {
	return svc.storer.NPCs(ctx, svc.gameServerId)
}
