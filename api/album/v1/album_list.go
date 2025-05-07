package v1

import (
	"taylor-music-back/internal/model/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAlbumListReq struct {
	g.Meta  `path:"/album/list" method:"get" tags:"Album" summary:"获取专辑列表"`
	Page    int    `v:"min:1#页码最小值为1" d:"1" dc:"页码"`
	Size    int    `v:"min:1#每页数量最小值为1" d:"10" dc:"每页数量"`
	Keyword string `d:"" dc:"搜索关键字(专辑标题)"`
}

type GetAlbumListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Data   struct {
		Records []dto.AlbumListDTO `json:"records"`
		Total   int                `json:"total"`
		Size    int                `json:"size"`
		Current int                `json:"current"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
