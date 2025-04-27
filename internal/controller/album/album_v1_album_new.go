package album

import (
	"taylor-music-back/api/album"
)

type ControllerV1 struct{}

func NewV1() album.IAlbumV1 {
	return &ControllerV1{}
}
