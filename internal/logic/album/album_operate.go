package album

import (
	"context"
	v1 "taylor-music-back/api/album/v1"
	"taylor-music-back/internal/dao"
	"taylor-music-back/internal/model/dto"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

func (s *sAlbum) CreateAlbum(ctx context.Context, v1 *v1.CreateAlbumReq) error {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:CreateAlbum | 创建专辑")
	if _, err := dao.Albums.Ctx(ctx).Data(g.Map{
		"AlbumUuid":       uuid.New().String(),
		"Title":           v1.Title,
		"ReleaseDate":     v1.ReleaseDate,
		"CoverImage":      v1.CoverImage,
		"BackgroundStory": v1.BackgroundStory,
		"Description":     v1.Description,
		"Producer":        v1.Producer,
		"TotalSongs":      0,
		"TotalDuration":   "0:00",
	}).Save(); err != nil {
		return err
	}
	return nil
}

func (s *sAlbum) EditAlbum(ctx context.Context, v1 *v1.EditAlbumReq) error {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:EditAlbum | 编辑专辑")
	if _, err := dao.Albums.Ctx(ctx).Data(g.Map{
		"AlbumUuid":       v1.AlbumUuid,
		"Title":           v1.Title,
		"ReleaseDate":     v1.ReleaseDate,
		"CoverImage":      v1.CoverImage,
		"BackgroundStory": v1.BackgroundStory,
		"Description":     v1.Description,
		"Producer":        v1.Producer,
	}).Where("album_uuid = ?", v1.AlbumUuid).Update(); err != nil {
		return err
	}
	return nil
}

func (s *sAlbum) DeleteAlbum(ctx context.Context, albumUuid string) error {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:DeleteAlbum | 删除专辑")

	// 先删除与专辑关联的歌曲
	if _, err := dao.Songs.Ctx(ctx).Where("album_id = ?", albumUuid).Delete(); err != nil {
		return err
	}

	// 然后删除专辑
	if _, err := dao.Albums.Ctx(ctx).Where("album_uuid = ?", albumUuid).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *sAlbum) GetAlbumOne(ctx context.Context, albumUuid string) (getAlbumOne dto.AlbumListDTO, err error) {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:GetAlbumOne | 获取专辑详情")
	album, err := dao.Albums.Ctx(ctx).Where("album_uuid = ?", albumUuid).One()
	if err != nil {
		return dto.AlbumListDTO{}, err
	}
	getAlbumOne = dto.AlbumListDTO{
		AlbumUuid:       album["album_uuid"].String(),
		Title:           album["title"].String(),
		ReleaseDate:     gtime.NewFromTime(album["release_date"].Time()),
		TotalSongs:      album["total_songs"].Int(),
		TotalDuration:   album["total_duration"].String(),
		CoverImage:      album["cover_image"].String(),
		Description:     album["description"].String(),
		BackgroundStory: album["background_story"].String(),
		Producer:        album["producer"].String(),
		CreatedAt:       gtime.NewFromTime(album["created_at"].Time()),
		UpdatedAt:       gtime.NewFromTime(album["updated_at"].Time()),
	}
	return getAlbumOne, nil
}

// GetAlbumSongCount 获取专辑中的歌曲数量
func (s *sAlbum) GetAlbumSongCount(ctx context.Context, albumUuid string) (int, error) {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:GetAlbumSongCount | 获取专辑歌曲数量")
	count, err := dao.Songs.Ctx(ctx).Where("album_id = ?", albumUuid).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// IncrementAlbumSongCount 增加专辑歌曲数量
func (s *sAlbum) IncrementAlbumSongCount(ctx context.Context, albumUuid string) error {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:IncrementAlbumSongCount | 增加专辑歌曲数量")

	// 获取当前的歌曲数量
	currentCount, err := s.GetAlbumSongCount(ctx, albumUuid)
	if err != nil {
		return err
	}

	// 更新专辑的歌曲数量
	if _, err := dao.Albums.Ctx(ctx).
		Data(g.Map{"total_songs": currentCount}).
		Where("album_uuid = ?", albumUuid).
		Update(); err != nil {
		return err
	}

	return nil
}

// DecrementAlbumSongCount 减少专辑歌曲数量
func (s *sAlbum) DecrementAlbumSongCount(ctx context.Context, albumUuid string) error {
	g.Log().Notice(ctx, "[LOGIC] AlbumLogic:DecrementAlbumSongCount | 减少专辑歌曲数量")

	// 获取当前的歌曲数量
	currentCount, err := s.GetAlbumSongCount(ctx, albumUuid)
	if err != nil {
		return err
	}

	// 更新专辑的歌曲数量
	if _, err := dao.Albums.Ctx(ctx).
		Data(g.Map{"total_songs": currentCount}).
		Where("album_uuid = ?", albumUuid).
		Update(); err != nil {
		return err
	}

	return nil
}
