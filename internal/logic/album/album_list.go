package album

import (
	"context"
	"fmt"
	"taylor-music-back/internal/dao"
	"taylor-music-back/internal/model/dto"
	"taylor-music-back/internal/model/entity"

	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取专辑列表
func (s *sAlbum) GetAlbumList(ctx context.Context, page int, size int, keyword string) (getAlbumList []dto.AlbumListDTO, totalCount int, err error) {
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

	// 获取总记录数
	totalCount, err = query.Count()
	if err != nil {
		return nil, 0, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}

	// 分页查询
	if err := query.Page(page, size).Scan(&albums); err != nil {
		return nil, 0, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}

	// 将 albums 转换为 dto.AlbumListDTO
	for _, album := range albums {
		// 获取专辑下所有歌曲并计算总时长
		var songs []entity.Songs
		err := dao.Songs.Ctx(ctx).Where("album_id", album.AlbumUuid).Scan(&songs)
		if err != nil {
			return nil, 0, berror.NewErrorHasError(bcode.ServerInternalError, err)
		}

		// 计算总时长（秒）
		var totalDurationSeconds int
		for _, song := range songs {
			totalDurationSeconds += song.Duration
		}

		// 转换为小时:分钟:秒格式
		hours := totalDurationSeconds / 3600
		minutes := (totalDurationSeconds % 3600) / 60
		seconds := totalDurationSeconds % 60
		formattedDuration := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

		getAlbumList = append(getAlbumList, dto.AlbumListDTO{
			AlbumUuid:       album.AlbumUuid,
			Title:           album.Title,
			ReleaseDate:     album.ReleaseDate,
			TotalSongs:      album.TotalSongs,
			TotalDuration:   formattedDuration,
			CoverImage:      album.CoverImage,
			Description:     album.Description,
			BackgroundStory: album.BackgroundStory,
			Producer:        album.Producer,
			CreatedAt:       album.CreatedAt,
			UpdatedAt:       album.UpdatedAt,
		})
	}

	// 如果没有记录返回空数组而不是错误
	if len(albums) == 0 {
		return []dto.AlbumListDTO{}, totalCount, nil
	}

	return getAlbumList, totalCount, nil
}
