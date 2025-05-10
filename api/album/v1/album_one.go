package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetAlbumOneReq struct {
	g.Meta    `path:"/album/:album_uuid" method:"get" tags:"Album" summary:"获取专辑详情"`
	AlbumUuid string `json:"album_uuid" in:"path" v:"required#专辑ID不能为空"`
}

type GetAlbumOneRes struct {
	g.Meta          `mime:"application/json" example:"json"`
	Title           string `json:"title" example:"专辑名称"			`
	ReleaseDate     string `json:"release_date" example:"发行日期"`
	TotalSongs      int    `json:"total_songs" example:"歌曲总数"`
	TotalDuration   string `json:"total_duration" example:"总共时长"`
	Producer        string `json:"producer" example:"唱片公司"`
	CoverImage      string `json:"cover_image" example:"专辑封面"`
	Description     string `json:"description" example:"专辑描述"`
	BackgroundStory string `json:"background_story" example:"背景故事"`
	CreatedAt       string `json:"created_at" example:"创建时间"`
	UpdatedAt       string `json:"updated_at" example:"更新时间"`
}
