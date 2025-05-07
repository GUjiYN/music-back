package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EditSongReq struct {
	g.Meta         `path:"/song/edit" method:"put" tags:"Song" summary:"编辑歌曲"`
	SongUuid       string `json:"song_uuid" in:"query" v:"required#歌曲ID不能为空" dc:"歌曲ID"`
	SongTitle      string `json:"song_title" v:"required#歌曲标题不能为空" dc:"歌曲标题"`
	AlbumId        string `json:"album_id" v:"required#专辑ID不能为空" dc:"专辑ID"`
	Duration       string `json:"duration" v:"required#时长不能为空" dc:"时长"`
	Lyrics         string `json:"lyrics" v:"required#歌词不能为空" dc:"歌词"`
	Writer         string `json:"writer" v:"required#作词不能为空" dc:"作词"`
	Producer       string `json:"producer" v:"required#制作人不能为空" dc:"制作人"`
	IsSingle       bool   `json:"is_single" v:"required#是否单曲不能为空" dc:"是否单曲"`
	ReleaseDate    string `json:"release_date" v:"required#发行日期不能为空" dc:"发行日期"`
	ReleaseVersion string `json:"release_version" v:"required#发行版本不能为空" dc:"发行版本"`
	Genre          string `json:"genre" v:"required#曲风不能为空" dc:"曲风"`
	Label          string `json:"label" v:"required#标签不能为空" dc:"标签"`
	Language       string `json:"language" v:"required#语种不能为空" dc:"语种"`
	Instruments    string `json:"instruments" v:"required#乐器不能为空" dc:"乐器"`
}

type EditSongRes struct {
	g.Meta  `mime:"application/json" example:"json"`
	Message string `json:"message" dc:"消息"`
	Code    int    `json:"code" dc:"状态码"`
}
