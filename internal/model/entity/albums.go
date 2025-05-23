// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Albums is the golang structure for table albums.
type Albums struct {
	AlbumUuid       string      `json:"albumUuid"       orm:"album_uuid"       description:"专辑表主键"` // 专辑表主键
	Title           string      `json:"title"           orm:"title"            description:"专辑名称"`  // 专辑名称
	ReleaseDate     *gtime.Time `json:"releaseDate"     orm:"release_date"     description:"发行日期"`  // 发行日期
	TotalSongs      int         `json:"totalSongs"      orm:"total_songs"      description:"歌曲总数"`  // 歌曲总数
	TotalDuration   string      `json:"totalDuration"   orm:"total_duration"   description:"总共时长"`  // 总共时长
	CoverImage      string      `json:"coverImage"      orm:"cover_image"      description:"专辑封面"`  // 专辑封面
	Description     string      `json:"description"     orm:"description"      description:"专辑描述"`  // 专辑描述
	BackgroundStory string      `json:"backgroundStory" orm:"background_story" description:"背景故事"`  // 背景故事
	Producer        string      `json:"producer"        orm:"producer"         description:"唱片公司"`  // 唱片公司
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`  // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`  // 更新时间
}
