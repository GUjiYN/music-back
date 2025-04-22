package api

import (
	"context"

	"taylor-music-back/api/request"
)

type AlbumApiV1 interface {
	List(ctx context.Context, req *request.AlbumListReq) (res *request.AlbumListRes, err error)
	Create(ctx context.Context, req *request.AlbumCreateReq) (res *request.AlbumCreateRes, err error)
	Edit(ctx context.Context, req *request.AlbumEditReq) (res *request.AlbumEditRes, err error)
	Delete(ctx context.Context, req *request.AlbumDeleteReq) (res *request.AlbumDeleteRes, err error)
}

type SongApiV1 interface {
	Create(ctx context.Context, req *request.SongCreateReq) (res *request.SongCreateRes, err error)
}
