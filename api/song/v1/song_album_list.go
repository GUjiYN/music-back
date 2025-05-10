package v1

import (
	"taylor-music-back/internal/model/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetSongsByAlbumReq struct {
	g.Meta  `path:"/song/album/list" method:"get" tags:"Song" summary:"根据专辑ID获取歌曲列表"`
	AlbumId string `v:"required#专辑ID不能为空" dc:"专辑ID"`
	Page    int    `v:"min:1#页码最小值为1" d:"1" dc:"页码"`
	Size    int    `v:"min:1#每页数量最小值为1" d:"10" dc:"每页数量"`
}

type GetSongsByAlbumRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Data   struct {
		Records []dto.SongListDTO `json:"records"`
		Total   int               `json:"total"`
		Size    int               `json:"size"`
		Current int               `json:"current"`
	} `json:"data"`
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Output       interface{} `json:"output"`
	ErrorMessage string      `json:"error_message"`
}
