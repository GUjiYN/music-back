// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AlbumsDao is the data access object for the table albums.
type AlbumsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AlbumsColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AlbumsColumns defines and stores column names for the table albums.
type AlbumsColumns struct {
	AlbumUuid       string // 专辑表主键
	Title           string // 专辑名称
	ReleaseDate     string // 发行日期
	TotalSongs      string // 歌曲总数
	TotalDuration   string // 总共时长
	CoverImage      string // 专辑封面
	Description     string // 专辑描述
	BackgroundStory string // 背景故事
	Producer        string // 唱片公司
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// albumsColumns holds the columns for the table albums.
var albumsColumns = AlbumsColumns{
	AlbumUuid:       "album_uuid",
	Title:           "title",
	ReleaseDate:     "release_date",
	TotalSongs:      "total_songs",
	TotalDuration:   "total_duration",
	CoverImage:      "cover_image",
	Description:     "description",
	BackgroundStory: "background_story",
	Producer:        "producer",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewAlbumsDao creates and returns a new DAO object for table data access.
func NewAlbumsDao(handlers ...gdb.ModelHandler) *AlbumsDao {
	return &AlbumsDao{
		group:    "default",
		table:    "albums",
		columns:  albumsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AlbumsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AlbumsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AlbumsDao) Columns() AlbumsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AlbumsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AlbumsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AlbumsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
