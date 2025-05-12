// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Albums is the golang structure of table albums for DAO operations like Where/Data.
type Albums struct {
	g.Meta          `orm:"table:albums, do:true"`
	AlbumUuid       interface{} // 专辑表主键
	Title           interface{} // 专辑名称
	ReleaseDate     *gtime.Time // 发行日期
	TotalSongs      interface{} // 歌曲总数
	TotalDuration   interface{} // 总共时长
	CoverImage      interface{} // 专辑封面
	Description     interface{} // 专辑描述
	BackgroundStory interface{} // 背景故事
	Producer        interface{} // 唱片公司
	SongsNumber     interface{} // 歌曲数目
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
