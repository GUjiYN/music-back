package album

import (
	"context"
	v1 "taylor-music-back/api/album/v1"

	"taylor-music-back/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
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
	if _, err := dao.Albums.Ctx(ctx).Where("album_uuid = ?", albumUuid).Delete(); err != nil {
		return err
	}
	return nil
}
