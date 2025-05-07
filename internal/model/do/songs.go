// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Songs is the golang structure of table songs for DAO operations like Where/Data.
type Songs struct {
	g.Meta         `orm:"table:songs, do:true"`
	SongUuid       interface{} // 歌曲表主键
	AlbumId        interface{} // 专辑id
	SongTitle      interface{} // 歌曲标题
	Duration       interface{} // 歌曲时长
	Lyrics         interface{} // 歌词
	Writer         interface{} // 作词人
	Producer       interface{} // 作曲人
	IsSingle       interface{} // 是否为单曲发行
	Genre          interface{} // 曲风
	Label          interface{} // 标签
	Language       interface{} // 语种
	Instruments    interface{} // 乐器
	ReleaseDate    *gtime.Time // 歌曲发行日期
	ReleaseVersion interface{} // 发行版本
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
