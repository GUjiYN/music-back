// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package song

import (
	"context"

	"taylor-music-back/api/song/v1"
)

type ISongV1 interface {
	GetSongList(ctx context.Context, req *v1.GetSongListReq) (res *v1.GetSongListRes, err error)
	CreateSong(ctx context.Context, req *v1.CreateSongReq) (res *v1.CreateSongRes, err error)
	EditSong(ctx context.Context, req *v1.EditSongReq) (res *v1.EditSongRes, err error)
	DeleteSong(ctx context.Context, req *v1.DeleteSongReq) (res *v1.DeleteSongRes, err error)
}
