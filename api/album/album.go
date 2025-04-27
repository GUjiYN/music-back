package album

import (
	"context"

	"taylor-music-back/api/album/v1"
)

type IAlbumV1 interface {
	GetAlbumList(ctx context.Context, req *v1.GetAlbumListReq) (res *v1.GetAlbumListRes, err error)
	CreateAlbum(ctx context.Context, req *v1.CreateAlbumRequest) (res *v1.CreateAlbumResponse, err error)
	EditAlbum(ctx context.Context, req *v1.EditAlbumRequest) (res *v1.EditAlbumResponse, err error)
}
