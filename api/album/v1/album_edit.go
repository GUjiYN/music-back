package v1

import "github.com/gogf/gf/v2/frame/g"

type EditAlbumReq struct {
	g.Meta          `path:"/album/edit" method:"put" tags:"Album" summary:"编辑专辑"`
	AlbumId         string `json:"albumId" in:"query" v:"required#专辑ID不能为空" dc:"专辑ID"`
	Title           string `json:"title" v:"required#标题不能为空" dc:"标题"`
	ReleaseDate     string `json:"releaseDate" v:"required#发行日期不能为空" dc:"发行日期"`
	CoverImage      string `json:"coverImage" v:"required#封面不能为空" dc:"封面"`
	Description     string `json:"description" v:"required#描述不能为空" dc:"描述"`
	BackgroundStory string `json:"backgroundStory" v:"required#背景故事不能为空" dc:"背景故事"`
	Producer        string `json:"producer" v:"required#唱片公司不能为空" dc:"唱片公司"`
}

type EditAlbumRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
