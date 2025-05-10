// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"taylor-music-back/api/song/v1"
	"taylor-music-back/internal/model/dto"
)

type (
	ISong interface {
		GetSongList(ctx context.Context, page int, size int) (getSongList []dto.SongListDTO, err error)
		GetSongsByAlbum(ctx context.Context, albumId string, page int, size int) (getSongList []dto.SongListDTO, err error)
		CreateSong(ctx context.Context, req *v1.CreateSongReq) error
		EditSong(ctx context.Context, req *v1.EditSongReq) error
		DeleteSong(ctx context.Context, req *v1.DeleteSongReq) error
	}
)

var (
	localSong ISong
)

func Song() ISong {
	if localSong == nil {
		panic("implement not found for interface ISong, forgot register?")
	}
	return localSong
}

func RegisterSong(i ISong) {
	localSong = i
}
