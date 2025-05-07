package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSongOneReq struct {
	g.Meta   `path:"/song/one" method:"get" tags:"Song" summary:"获取歌曲详情"`
	SongUuid string `json:"song_uuid" in:"query" v:"required#歌曲ID不能为空"`
}

type GetSongOneRes struct {
	g.Meta    `mime:"application/json" example:"json"`
	SongUuid  string `json:"song_uuid" example:"歌曲ID"`
	SongTitle string `json:"song_title" example:"歌曲标题"`
	AlbumUuid string `json:"album_uuid" example:"专辑ID"`
}
