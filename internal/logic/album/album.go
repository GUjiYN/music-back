package album

import (
	"taylor-music-back/internal/service"
)

type sAlbum struct{}

func init() {
	service.RegisterAlbum(New())
}

func New() *sAlbum {
	return &sAlbum{}
}
