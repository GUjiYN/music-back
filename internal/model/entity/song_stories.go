// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SongStories is the golang structure for table song_stories.
type SongStories struct {
	StoryUuid string      `json:"storyUuid" orm:"story_uuid" description:"创作故事表主键"` // 创作故事表主键
	SongId    string      `json:"songId"    orm:"song_id"    description:"歌曲id"`    // 歌曲id
	Content   string      `json:"content"   orm:"content"    description:"具体内容"`    // 具体内容
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`    // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`    // 更新时间
}
