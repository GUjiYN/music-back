// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package album

import (
	"context"

	"taylor-music-back/api/album/v1"
)

type IAlbumV1 interface {
	GetAlbumList(ctx context.Context, req *v1.GetAlbumListReq) (res *v1.GetAlbumListRes, err error)
}
