package albumServiceImpl

import (
	"taylor-music-back/utility/ErrorCode"
	"taylor-music-back/utility/ResultUtil"

	"taylor-music-back/internal/model/do"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"taylor-music-back/api/request"
	"taylor-music-back/internal/model/entity"
	"taylor-music-back/internal/dao"
)

type DefaultAlbumImpl struct{}

// GetList 获取专辑列表
func (l *DefaultAlbumImpl) GetAlbumList(req *ghttp.Request) {
	albumDO := dao.Albums.GetAlbumList()
	if albumDO != nil {
		ResultUtil.Success(req, "查询成功", albumDO)
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
	}
}

// CreateAlbum 创建专辑
func (l *DefaultAlbumImpl) CreateAlbum(req *ghttp.Request, album *request.AlbumCreateReq) {
	albumDO := entity.Albums{
		AlbumUuid:       uuid.New().String(),
		Title:           album.Title,
		ReleaseDate:     gtime.NewFromStr(album.ReleaseDate),
		CoverImage:      album.CoverImage,
		Description:     album.Description,
		BackgroundStory: album.BackgroundStory,
		Producer:        album.Producer,
		CreatedAt:       gtime.Now(),
		UpdatedAt:       gtime.Now(),
	}
	if dao.Albums.Create(albumDO) {
		ResultUtil.Success(req, "创建成功", albumDO)
	} else {
		ResultUtil.Error(req, ErrorCode.ServerDatabaseInteriorError, "创建失败")
	}
}

// UpdateAlbum 更新专辑
func (l *DefaultAlbumImpl) UpdateAlbum(req *ghttp.Request, album entity.AlbumEditVO) {
	albumDO := do.AlbumDO{
		AlbumUuid:       album.AlbumUuid,
		Title:           album.Title,
		ReleaseDate:     album.ReleaseDate,
		CoverImage:      album.CoverImage,
		Description:     album.Description,
		BackgroundStory: album.BackgroundStory,
		Producer:        album.Producer,
		UpdatedAt:       gtime.Now(),
	}
	if albumDAO.UpdateAlbum(albumDO) {
		ResultUtil.Success(req, "更新成功", album)
	} else {
		ResultUtil.Error(req, ErrorCode.ServerDatabaseInteriorError, "更新失败")
	}
}


// DeleteAlbum 删除专辑
func (l *DefaultAlbumImpl) DeleteAlbum(req *ghttp.Request, album entity.AlbumDeleteVO) {
	if albumDAO.DeleteAlbum(album.AlbumUuid) {
		ResultUtil.Success(req, "删除成功", album)
	} else {
		ResultUtil.Error(req, ErrorCode.ServerDatabaseInteriorError, "删除失败")
	}
}
