package album

import (
	"context"
	"taylor-music-back/internal/dao"
	"taylor-music-back/internal/model/dto"
	"taylor-music-back/internal/model/entity"

	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取专辑列表
func (s *sAlbum) GetAlbumList(ctx context.Context, page int, size int, keyword string) (getAlbumList []dto.AlbumListDTO, err error) {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:GetAlbumList | 获取专辑列表")
	// 从数据库获取专辑列表
	var albums []entity.Albums

	// 创建查询构建器
	query := dao.Albums.Ctx(ctx)

	// 如果关键字不为空，添加模糊查询条件
	if keyword != "" {
		query = query.WhereLike("title", "%"+keyword+"%")
		g.Log().Notice(ctx, "[LOGIC] AlbumLogic:GetAlbumList | 使用关键字搜索: %s", keyword)
	}

	// 分页查询
	if err := query.Page(page, size).Scan(&albums); err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}

	// 将 albums 转换为 dto.AlbumListDTO
	for _, album := range albums {
		getAlbumList = append(getAlbumList, dto.AlbumListDTO{
			AlbumUuid:       album.AlbumUuid,
			Title:           album.Title,
			ReleaseDate:     album.ReleaseDate,
			TotalSongs:      album.TotalSongs,
			TotalDuration:   album.TotalDuration,
			CoverImage:      album.CoverImage,
			Description:     album.Description,
			BackgroundStory: album.BackgroundStory,
			Producer:        album.Producer,
			CreatedAt:       album.CreatedAt,
			UpdatedAt:       album.UpdatedAt,
		})
	}

	// 检查是否获取到专辑
	if len(albums) == 0 {
		return nil, berror.NewError(bcode.NotExist, "未查询到专辑")
	}

	return getAlbumList, nil
}
