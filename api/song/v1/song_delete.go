package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteSongReq struct {
	g.Meta   `path:"/song/delete" method:"post" tags:"Song" summary:"删除歌曲"`
	SongUuid string `json:"songUuid" v:"required#歌曲ID不能为空" dc:"歌曲ID"`
}

type DeleteSongRes struct {
	g.Meta  `mime:"application/json" example:"json"`
	Message string `json:"message" dc:"消息"`
	Code    int    `json:"code" dc:"状态码"`
}
