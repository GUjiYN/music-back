// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "taylor-music-back/api/album/v1"
	"taylor-music-back/internal/model/dto"
)

type (
	IAlbum interface {
		// 获取专辑列表
		GetAlbumList(ctx context.Context, page int, size int, keyword string) (getAlbumList []dto.AlbumListDTO, totalCount int, err error)
		CreateAlbum(ctx context.Context, v1 *v1.CreateAlbumReq) error
		EditAlbum(ctx context.Context, v1 *v1.EditAlbumReq) error
		DeleteAlbum(ctx context.Context, albumUuid string) error
		GetAlbumOne(ctx context.Context, albumUuid string) (getAlbumOne dto.AlbumListDTO, err error)
		// 获取专辑中的歌曲数量
		GetAlbumSongCount(ctx context.Context, albumUuid string) (int, error)
		// 增加专辑歌曲数量
		IncrementAlbumSongCount(ctx context.Context, albumUuid string) error
		// 减少专辑歌曲数量
		DecrementAlbumSongCount(ctx context.Context, albumUuid string) error
	}
)

var (
	localAlbum IAlbum
)

func Album() IAlbum {
	if localAlbum == nil {
		panic("implement not found for interface IAlbum, forgot register?")
	}
	return localAlbum
}

func RegisterAlbum(i IAlbum) {
	localAlbum = i
}
