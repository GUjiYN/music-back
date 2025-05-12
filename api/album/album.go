package album

import (
	"context"

	v1 "taylor-music-back/api/album/v1"
)

type IAlbumV1 interface {
	GetAlbumList(ctx context.Context, req *v1.GetAlbumListReq) (res *v1.GetAlbumListRes, err error)
	CreateAlbum(ctx context.Context, req *v1.CreateAlbumReq) (res *v1.CreateAlbumRes, err error)
	EditAlbum(ctx context.Context, req *v1.EditAlbumReq) (res *v1.EditAlbumRes, err error)
	DeleteAlbum(ctx context.Context, req *v1.DeleteAlbumReq) (res *v1.DeleteAlbumRes, err error)
	GetAlbumOne(ctx context.Context, req *v1.GetAlbumOneReq) (res *v1.GetAlbumOneRes, err error)
}
