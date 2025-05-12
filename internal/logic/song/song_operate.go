package song

import (
	"context"
	v1 "taylor-music-back/api/song/v1"

	"taylor-music-back/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

func (s *sSong) CreateSong(ctx context.Context, v1 *v1.CreateSongReq) error {
	g.Log().Notice(ctx, "[LOGIC] SongLogic:CreateSong | 创建歌曲")
	if _, err := dao.Songs.Ctx(ctx).Data(g.Map{
		"SongUuid":       uuid.New().String(),
		"SongTitle":      v1.SongTitle,
		"AlbumId":        v1.AlbumId,
		"Duration":       v1.Duration,
		"Lyrics":         v1.Lyrics,
		"Writer":         v1.Writer,
		"Producer":       v1.Producer,
		"IsSingle":       v1.IsSingle,
		"ReleaseDate":    v1.ReleaseDate,
		"ReleaseVersion": v1.ReleaseVersion,
		"Language":       v1.Language,
		"Genre":          v1.Genre,
		"Label":          v1.Label,
		"Instruments":    v1.Instruments,
	}).Save(); err != nil {
		return err
	}
	return nil
}

func (s *sSong) EditSong(ctx context.Context, v1 *v1.EditSongReq) error {
	g.Log().Notice(ctx, "[LOGIC] SongLogic:EditSong | 编辑歌曲")
	if _, err := dao.Songs.Ctx(ctx).Data(g.Map{
		"SongUuid":    v1.SongUuid,
		"SongTitle":   v1.SongTitle,
		"AlbumId":     v1.AlbumId,
		"Duration":    v1.Duration,
		"Lyrics":      v1.Lyrics,
		"Writer":      v1.Writer,
		"Producer":    v1.Producer,
		"IsSingle":    v1.IsSingle,
		"ReleaseDate": v1.ReleaseDate,
	}).Where("Song_uuid = ?", v1.SongUuid).Update(); err != nil {
		return err
	}
	return nil
}

func (s *sSong) DeleteSong(ctx context.Context, v1 *v1.DeleteSongReq) error {
	g.Log().Notice(ctx, "[LOGIC] SongLogic:DeleteSong | 删除歌曲")
	if _, err := dao.Songs.Ctx(ctx).Where("Song_uuid = ?", v1.SongUuid).Delete(); err != nil {
		return err
	}
	return nil
}
