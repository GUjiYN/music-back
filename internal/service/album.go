// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"taylor-music-back/internal/model/dto"
)

type (
	IAlbum interface {
		// 获取专辑列表
		GetAlbumList(ctx context.Context, page int, size int) (getAlbumList []dto.AlbumListDTO, err error)
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
