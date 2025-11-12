package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
)

func (svc Service) NPCResponses(ctx context.Context) (map[int]retro.NPCResponse, error) {
	return svc.storer.NPCResponses(ctx)
}
