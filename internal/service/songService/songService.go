package songService

import (
	"context"
	"taylor-music-back/internal/model/entity"
)

type SongService interface {
	GetList(ctx context.Context, page, size int) (list []*entity.Song, total int, pages int, err error)
	GetOne(ctx context.Context, songUuid string) (song *entity.Song, err error)
	GetByNumber(ctx context.Context, number string) (song *entity.Song, err error)
	Create(ctx context.Context, song *entity.Song) (songUuid string, err error)
}

