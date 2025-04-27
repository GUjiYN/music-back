package song

import (
	"context"
	"taylor-music-back/api/song/v1"
	"taylor-music-back/internal/service"
)

func (c *ControllerV1) EditSong(ctx context.Context, req *v1.EditSongReq) (res *v1.EditSongRes, err error) {
	err = service.Song().EditSong(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.EditSongRes{
		Message: "编辑歌曲成功",
		Code:    200,
	}, nil
}
