// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SongsDao is the data access object for the table songs.
type SongsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SongsColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SongsColumns defines and stores column names for the table songs.
type SongsColumns struct {
	SongUuid       string // 歌曲表主键
	AlbumId        string // 专辑id
	SongTitle      string // 歌曲标题
	Duration       string // 歌曲时长
	Lyrics         string // 歌词
	Writer         string // 作词人
	Producer       string // 作曲人
	IsSingle       string // 是否为单曲发行
	Genre          string // 曲风
	Label          string // 标签
	Language       string // 语种
	Instruments    string // 乐器
	ReleaseDate    string // 歌曲发行日期
	ReleaseVersion string // 发行版本
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// songsColumns holds the columns for the table songs.
var songsColumns = SongsColumns{
	SongUuid:       "song_uuid",
	AlbumId:        "album_id",
	SongTitle:      "song_title",
	Duration:       "duration",
	Lyrics:         "lyrics",
	Writer:         "writer",
	Producer:       "producer",
	IsSingle:       "is_single",
	Genre:          "genre",
	Label:          "label",
	Language:       "language",
	Instruments:    "instruments",
	ReleaseDate:    "release_date",
	ReleaseVersion: "release_version",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewSongsDao creates and returns a new DAO object for table data access.
func NewSongsDao(handlers ...gdb.ModelHandler) *SongsDao {
	return &SongsDao{
		group:    "default",
		table:    "songs",
		columns:  songsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SongsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SongsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SongsDao) Columns() SongsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SongsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SongsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SongsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
