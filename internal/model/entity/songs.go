// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Songs is the golang structure for table songs.
type Songs struct {
	SongUuid       string      `json:"songUuid"       orm:"song_uuid"       description:"歌曲表主键"`   // 歌曲表主键
	AlbumId        string      `json:"albumId"        orm:"album_id"        description:"专辑id"`    // 专辑id
	SongTitle      string      `json:"songTitle"      orm:"song_title"      description:"歌曲标题"`    // 歌曲标题
	Duration       int         `json:"duration"       orm:"duration"        description:"歌曲时长"`    // 歌曲时长
	Lyrics         string      `json:"lyrics"         orm:"lyrics"          description:"歌词"`      // 歌词
	Writer         string      `json:"writer"         orm:"writer"          description:"作词人"`     // 作词人
	Producer       string      `json:"producer"       orm:"producer"        description:"作曲人"`     // 作曲人
	IsSingle       int         `json:"isSingle"       orm:"is_single"       description:"是否为单曲发行"` // 是否为单曲发行
	Genre          string      `json:"genre"          orm:"genre"           description:"曲风"`      // 曲风
	Label          string      `json:"label"          orm:"label"           description:"标签"`      // 标签
	Language       string      `json:"language"       orm:"language"        description:"语种"`      // 语种
	Instruments    string      `json:"instruments"    orm:"instruments"     description:"乐器"`      // 乐器
	ReleaseDate    *gtime.Time `json:"releaseDate"    orm:"release_date"    description:"歌曲发行日期"`  // 歌曲发行日期
	ReleaseVersion string      `json:"releaseVersion" orm:"release_version" description:"发行版本"`    // 发行版本
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`    // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`    // 更新时间
}
