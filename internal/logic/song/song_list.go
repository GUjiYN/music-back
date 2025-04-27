package song

import (
	"context"
	"taylor-music-back/internal/dao"
	"taylor-music-back/internal/model/dto"
	"taylor-music-back/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sSong) GetSongList(ctx context.Context, page int, size int) (getSongList []dto.SongListDTO, err error) {
	g.Log().Notice(ctx, "[LOGIC] SongLogic:GetSongList | 获取歌曲列表")

	var songs []entity.Songs
	if err := dao.Songs.Ctx(ctx).Page(page, size).Scan(&songs); err != nil {
		return nil, err
	}

	for _, song := range songs {
		getSongList = append(getSongList, dto.SongListDTO{
			SongUuid:    song.SongUuid,
			SongTitle:   song.SongTitle,
			AlbumId:     song.AlbumId,
			Duration:    song.Duration,
			Lyrics:      song.Lyrics,
			Writer:      song.Writer,
			Producer:    song.Producer,
			IsSingle:    song.IsSingle,
			ReleaseDate: song.ReleaseDate,
			CreatedAt:   song.CreatedAt,
			UpdatedAt:   song.UpdatedAt,
		})
	}

	return getSongList, nil
}
