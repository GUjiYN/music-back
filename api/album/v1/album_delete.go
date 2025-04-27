package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteAlbumRequest struct {
	g.Meta    `path:"/album/delete" method:"delete" tags:"Album" summary:"删除专辑"`
	AlbumUuid string `json:"albumUuid" in:"query" v:"required#专辑ID不能为空" dc:"专辑ID"`
}

type DeleteAlbumResponse struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
