package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"taylor-music-back/internal/model/dto"
)

type GetSongListReq struct {
	g.Meta `path:"/song/list" method:"get" tags:"Song" summary:"获取歌曲列表"`
	Page   int `v:"min:1#页码最小值为1" d:"1" dc:"页码"`
	Size   int `v:"min:1#每页数量最小值为1" d:"10" dc:"每页数量"`
}

type GetSongListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Page   int               `json:"page"`
	Pages  int               `json:"pages"`
	Size   int               `json:"size"`
	Total  int               `json:"total"`
	List   []dto.SongListDTO `json:"list"`
}
