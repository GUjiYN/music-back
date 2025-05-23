package album

import (
	"context"
	"fmt"
	v1 "taylor-music-back/api/album/v1"
	"taylor-music-back/internal/dao"
	"taylor-music-back/internal/model/dto"
	"taylor-music-back/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

// 计算专辑总播放时间并格式化为小时:分钟:秒
func (s *sAlbum) calculateAlbumTotalDuration(ctx context.Context, albumUuid string) (string, error) {
	var songs []entity.Songs
	err := dao.Songs.Ctx(ctx).Where("album_id", albumUuid).Scan(&songs)
	if err != nil {
		return "", err
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
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
}

// 更新专辑总播放时间
func (s *sAlbum) updateAlbumTotalDuration(ctx context.Context, albumUuid string) error {
	duration, err := s.calculateAlbumTotalDuration(ctx, albumUuid)
	if err != nil {
		return err
	}

	_, err = dao.Albums.Ctx(ctx).
		Data(g.Map{"total_duration": duration}).
		Where("album_uuid = ?", albumUuid).
		Update()
	return err
}

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
		"TotalDuration":   "00:00:00",
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

	// 更新总播放时间
	if err := s.updateAlbumTotalDuration(ctx, albumUuid); err != nil {
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

	// 更新总播放时间
	if err := s.updateAlbumTotalDuration(ctx, albumUuid); err != nil {
		return err
	}

	return nil
}
