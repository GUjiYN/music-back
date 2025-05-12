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
			SongUuid:       song.SongUuid,
			SongTitle:      song.SongTitle,
			AlbumId:        song.AlbumId,
			Duration:       secondsToTime(song.Duration),
			Lyrics:         song.Lyrics,
			Writer:         song.Writer,
			Producer:       song.Producer,
			IsSingle:       song.IsSingle,
			Genre:          song.Genre,
			Label:          song.Label,
			Language:       song.Language,
			Instruments:    song.Instruments,
			ReleaseDate:    song.ReleaseDate,
			ReleaseVersion: song.ReleaseVersion,
			CreatedAt:      song.CreatedAt,
			UpdatedAt:      song.UpdatedAt,
		})
	}

	return getSongList, nil
}

func (s *sSong) GetSongsByAlbum(ctx context.Context, albumId string, page int, size int) (getSongList []dto.SongListDTO, err error) {
	g.Log().Notice(ctx, "[LOGIC] SongLogic:GetSongsByAlbum | 根据专辑ID获取歌曲列表")

	var songs []entity.Songs
	if err := dao.Songs.Ctx(ctx).Where("album_id", albumId).Page(page, size).Scan(&songs); err != nil {
		return nil, err
	}

	for _, song := range songs {
		getSongList = append(getSongList, dto.SongListDTO{
			SongUuid:       song.SongUuid,
			SongTitle:      song.SongTitle,
			AlbumId:        song.AlbumId,
			Duration:       secondsToTime(song.Duration),
			Lyrics:         song.Lyrics,
			Writer:         song.Writer,
			Producer:       song.Producer,
			IsSingle:       song.IsSingle,
			Genre:          song.Genre,
			Label:          song.Label,
			Language:       song.Language,
			Instruments:    song.Instruments,
			ReleaseDate:    song.ReleaseDate,
			ReleaseVersion: song.ReleaseVersion,
			CreatedAt:      song.CreatedAt,
			UpdatedAt:      song.UpdatedAt,
		})
	}

	return getSongList, nil
}
