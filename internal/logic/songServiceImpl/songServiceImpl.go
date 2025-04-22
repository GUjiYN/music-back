package songServiceImpl

import (
	"context"
	"taylor-music-back/internal/dao"
	"taylor-music-back/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

type SongLogic struct{}

var Song = &SongLogic{}

func (l *SongLogic) Create(ctx context.Context, song *entity.Song) (songUuid string, err error) {
	song1 := &entity.Song{
		SongUuid:    guid.S(),
		AlbumUuid:   song.AlbumUuid,
		SongTitle:   song.SongTitle,
		Duration:    song.Duration,
		Lyrics:      song.Lyrics,
		Writer:      song.Writer,
		Producer:    song.Producer,
		IsSingle:    song.IsSingle,
		ReleaseDate: song.ReleaseDate,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
	}

	return dao.Song.Create(ctx, song1)
}
