package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateAlbumReq struct {
	g.Meta          `path:"/album/create" method:"post" tags:"Album" summary:"创建专辑"`
	Title           string `json:"title" v:"required#标题不能为空" dc:"标题"`
	ReleaseDate     string `json:"releaseDate" v:"required#发行日期不能为空" dc:"发行日期"`
	CoverImage      string `json:"coverImage" v:"required#封面不能为空" dc:"封面"`
	Description     string `json:"description" v:"required#描述不能为空" dc:"描述"`
	BackgroundStory string `json:"backgroundStory" v:"required#背景故事不能为空" dc:"背景故事"`
	Producer        string `json:"producer" v:"required#唱片公司不能为空" dc:"唱片公司"`
}

type CreateAlbumRes struct {
	g.Meta  `mime:"application/json" example:"json"`
	Message string `json:"message" dc:"消息"`
	Code    int    `json:"code" dc:"状态码"`
}
