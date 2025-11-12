package retrosvc

import (
	"context"

	retro "github.com/hadamrd/retrodata"
	"github.com/hadamrd/retrodata/retrotyp"
)

func (svc Service) Classes(ctx context.Context) (map[retrotyp.ClassId]retro.Class, error) {
	return svc.storer.Classes(ctx)
}
