package song

import (
	"context"
	"taylor-music-back/api/song/v1"
	"taylor-music-back/internal/service"
)

func (c *ControllerV1) CreateSong(ctx context.Context, req *v1.CreateSongReq) (res *v1.CreateSongRes, err error) {
	err = service.Song().CreateSong(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.CreateSongRes{
		Message: "创建歌曲成功",
		Code:    200,
	}, nil
}
