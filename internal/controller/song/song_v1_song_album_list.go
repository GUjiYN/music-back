package song

import (
	"context"
	"taylor-music-back/internal/model/dto"
	"taylor-music-back/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "taylor-music-back/api/song/v1"
)

func (c *ControllerV1) GetSongsByAlbum(ctx context.Context, req *v1.GetSongsByAlbumReq) (res *v1.GetSongsByAlbumRes, err error) {
	songList, err := service.Song().GetSongsByAlbum(ctx, req.AlbumId, req.Page, req.Size)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "获取专辑歌曲列表失败", err.Error())
	}
	return &v1.GetSongsByAlbumRes{
		Data: struct {
			Records []dto.SongListDTO `json:"records"`
			Total   int               `json:"total"`
			Size    int               `json:"size"`
			Current int               `json:"current"`
		}{
			Records: songList,
			Total:   len(songList),
			Size:    req.Size,
			Current: req.Page,
		},
	}, nil
}
