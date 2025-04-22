// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SongStories is the golang structure of table song_stories for DAO operations like Where/Data.
type SongStories struct {
	g.Meta    `orm:"table:song_stories, do:true"`
	StoryUuid interface{} // 创作故事表主键
	SongId    interface{} // 歌曲id
	Content   interface{} // 具体内容
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
