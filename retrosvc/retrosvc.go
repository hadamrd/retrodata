package retrosvc

import (
	"errors"

	retro "github.com/hadamrd/retrodata"
)

type Config struct {
	GameServerId int
	Storer       retro.Storer
}

type Service struct {
	gameServerId int
	storer       retro.Storer
}

func NewService(c Config) (*Service, error) {
	if c.Storer == nil {
		return nil, errors.New("nil storer")
	}
	svc := &Service{
		gameServerId: c.GameServerId,
		storer:       c.Storer,
	}
	return svc, nil
}
