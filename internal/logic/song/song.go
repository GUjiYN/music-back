package song

import (
	"taylor-music-back/internal/service"
)

type sSong struct{}

func init() {
	service.RegisterSong(New())
}

func New() *sSong {
	return &sSong{}
}
