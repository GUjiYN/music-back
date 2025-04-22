package request

import (
	"github.com/gogf/gf/v2/frame/g"
)


type SongCreateReq struct {
	g.Meta `path:"/song/create" tags:"Song" method:"post" summary:"创建歌曲接口"`
	SongUuid string `json:"song_uuid"`
	AlbumUuid string `json:"album_uuid"`
	SongTitle string `json:"song_title"`
	Duration  string `json:"duration"`
	Lyrics    string `json:"lyrics"`
	Writer    string `json:"writer"`
	Producer  string `json:"producer"`
	IsSingle  bool   `json:"is_single"`
	ReleaseDate string `json:"release_date"`
}

type SongCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Message  string `json:"message"`
}



