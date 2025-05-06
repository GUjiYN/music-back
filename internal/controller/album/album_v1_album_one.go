package album

import (
	"context"
	"taylor-music-back/api/album/v1"
	"taylor-music-back/internal/service"
)

func (c *ControllerV1) GetAlbumOne(ctx context.Context, req *v1.GetAlbumOneReq) (res *v1.GetAlbumOneRes, err error) {
	album, err := service.Album().GetAlbumOne(ctx, req.AlbumUuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetAlbumOneRes{
		Title:           album.Title,
		ReleaseDate:     album.ReleaseDate.String(),
		TotalSongs:      album.TotalSongs,
		TotalDuration:   album.TotalDuration,
		Producer:        album.Producer,
		CoverImage:      album.CoverImage,
		Description:     album.Description,
		BackgroundStory: album.BackgroundStory,
		CreatedAt:       album.CreatedAt.String(),
		UpdatedAt:       album.UpdatedAt.String(),
	}, nil
}
