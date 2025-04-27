package song

import (
	"context"
	"taylor-music-back/api/song/v1"
	"taylor-music-back/internal/service"
)

func (c *ControllerV1) DeleteSong(ctx context.Context, req *v1.DeleteSongReq) (res *v1.DeleteSongRes, err error) {
	err = service.Song().DeleteSong(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteSongRes{
		Message: "删除歌曲成功",
		Code:    200,
	}, nil
}
