package album

import (
	"context"
	v1 "taylor-music-back/api/album/v1"
	"taylor-music-back/internal/service"
)

func (c *ControllerV1) DeleteAlbum(ctx context.Context, req *v1.DeleteAlbumReq) (res *v1.DeleteAlbumRes, err error) {
	err = service.Album().DeleteAlbum(ctx, req.AlbumUuid)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteAlbumRes{
		Message: "删除成功",
		Code:    200,
	}, nil
}
