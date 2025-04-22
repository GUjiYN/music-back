package albumService

import (
	"taylor-music-back/internal/logic/albumServiceImpl"
	"taylor-music-back/internal/model/entity"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NewAlbumService() AlbumService {
	return &albumServiceImpl.DefaultAlbumImpl{}
}

type AlbumService interface {
	GetAlbumList(*ghttp.Request)
	CreateAlbum(*ghttp.Request, entity.AlbumVO)
	UpdateAlbum(*ghttp.Request, entity.AlbumEditVO)
	DeleteAlbum(*ghttp.Request, entity.AlbumDeleteVO)
}
