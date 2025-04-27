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
}
