package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateSongReq struct {
	g.Meta      `path:"/song/create" method:"post" tags:"Song" summary:"创建歌曲"`
	SongTitle   string `json:"songTitle" v:"required#标题不能为空" dc:"标题"`
	AlbumId     string `json:"albumId" v:"required#专辑不能为空" dc:"专辑"`
	Duration    string `json:"duration" v:"required#时长不能为空" dc:"时长"`
	Lyrics      string `json:"lyrics" v:"required#歌词不能为空" dc:"歌词"`
	Writer      string `json:"writer" v:"required#作词不能为空" dc:"作词"`
	Producer    string `json:"producer" v:"required#制作人不能为空" dc:"制作人"`
	IsSingle    bool   `json:"isSingle" v:"required#是否单曲不能为空" dc:"是否单曲"`
	ReleaseDate string `json:"releaseDate" v:"required#发行日期不能为空" dc:"发行日期"`
}

type CreateSongRes struct {
	g.Meta  `mime:"application/json" example:"json"`
	Message string `json:"message" dc:"消息"`
	Code    int    `json:"code" dc:"状态码"`
}
