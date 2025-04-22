// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SongStoriesDao is the data access object for the table song_stories.
type SongStoriesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SongStoriesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SongStoriesColumns defines and stores column names for the table song_stories.
type SongStoriesColumns struct {
	StoryUuid string // 创作故事表主键
	SongId    string // 歌曲id
	Content   string // 具体内容
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// songStoriesColumns holds the columns for the table song_stories.
var songStoriesColumns = SongStoriesColumns{
	StoryUuid: "story_uuid",
	SongId:    "song_id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSongStoriesDao creates and returns a new DAO object for table data access.
func NewSongStoriesDao(handlers ...gdb.ModelHandler) *SongStoriesDao {
	return &SongStoriesDao{
		group:    "default",
		table:    "song_stories",
		columns:  songStoriesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SongStoriesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SongStoriesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SongStoriesDao) Columns() SongStoriesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SongStoriesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SongStoriesDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SongStoriesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
