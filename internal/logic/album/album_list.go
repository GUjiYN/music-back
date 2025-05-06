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
func (s *sAlbum) GetAlbumList(ctx context.Context, page int, size int) (getAlbumList []dto.AlbumListDTO, err error) {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:GetAlbumList | 获取专辑列表")
	// 从数据库获取专辑列表
	var albums []entity.Albums
	if err := dao.Albums.Ctx(ctx).Page(page, size).Scan(&albums); err != nil {
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
