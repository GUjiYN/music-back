package request

import (
	"taylor-music-back/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// AlbumListReq 获取专辑列表请求
type AlbumListReq struct {
	g.Meta `path:"/albumServiceImpl/list" tags:"Album" method:"get" summary:"获取专辑列表接口"`
	Page   int `v:"min:1#页码最小值为1" d:"1" dc:"页码"`
	Size   int `v:"min:1#每页数量最小值为1" d:"10" dc:"每页数量"`
}

// AlbumListRes 获取专辑列表响应
type AlbumListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	List   []*entity.AlbumVO `json:"list"`  // 专辑列表
	Total  int               `json:"total"` // 总数量
	Page   int               `json:"page"`  // 当前页码
	Size   int               `json:"size"`  // 每页数量
	Pages  int               `json:"pages"` // 总页数
}

// AlbumGetReq 获取专辑详情请求
type AlbumGetReq struct {
	g.Meta    `path:"/albumServiceImpl/detail" tags:"Album" method:"get" summary:"获取专辑详情接口"`
	AlbumUuid string `v:"required#专辑UUID不能为空" dc:"专辑UUID"`
}

// AlbumGetRes 获取专辑详情响应
type AlbumGetRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*entity.AlbumVO
}

// AlbumByNumberReq 根据专辑编号获取专辑请求
type AlbumByNumberReq struct {
	g.Meta `path:"/albumServiceImpl/number" tags:"Album" method:"get" summary:"根据专辑编号获取专辑接口"`
	Number string `v:"required#专辑编号不能为空" dc:"专辑编号"`
}

// AlbumByNumberRes 根据专辑编号获取专辑响应
type AlbumByNumberRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*entity.AlbumVO
}

// AlbumCreateReq 创建专辑请求
type AlbumCreateReq struct {
	g.Meta          `path:"/albumServiceImpl/create" tags:"Album" method:"post" summary:"创建专辑接口"`
	Title           string `v:"required#专辑名称不能为空" dc:"专辑名称"`
	CoverImage      string `v:"required#专辑封面不能为空" dc:"专辑封面"`
	ReleaseDate     string `v:"required#发行日期不能为空" dc:"发行日期"`
	Description     string `dc:"专辑描述"`
	BackgroundStory string `dc:"专辑背景描述"`
	Producer        string `dc:"唱片公司"`
}

// AlbumCreateRes 创建专辑响应
type AlbumCreateRes struct {
	g.Meta    `mime:"application/json" example:"json"`
	AlbumUuid string `json:"album_uuid"` // 创建的专辑UUID
	Message   string `json:"message"`    // 响应消息
}

// AlbumEditReq 更新专辑请求
type AlbumEditReq struct {
	g.Meta          `path:"/albumServiceImpl/edit" tags:"Album" method:"put" summary:"更新专辑接口"`
	AlbumUuid       string `in:"query" v:"required#专辑UUID不能为空" dc:"专辑UUID"`
	Title           string `v:"required#专辑名称不能为空" dc:"专辑名称"`
	CoverImage      string `v:"required#专辑封面不能为空" dc:"专辑封面"`
	ReleaseDate     string `v:"required#发行日期不能为空" dc:"发行日期"`
	Description     string `dc:"专辑描述"`
	BackgroundStory string `dc:"专辑背景描述"`
	Producer        string `dc:"唱片公司"`
}

// AlbumEditRes 更新专辑响应
type AlbumEditRes struct {
	g.Meta  `mime:"application/json" example:"json"`
	Message string `json:"message"` // 响应消息
}

// AlbumDeleteReq 删除专辑请求
type AlbumDeleteReq struct {
	g.Meta    `path:"/albumServiceImpl/delete" tags:"Album" method:"delete" summary:"删除专辑接口"`
	AlbumUuid string `v:"required#专辑UUID不能为空" dc:"专辑UUID"`
}

// AlbumDeleteRes 删除专辑响应
type AlbumDeleteRes struct {
	g.Meta  `mime:"application/json" example:"json"`
	Message string `json:"message"` // 响应消息
}
