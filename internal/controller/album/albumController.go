package album

import (
	"context"

	"taylor-music-back/api"
	"taylor-music-back/api/request"
	"taylor-music-back/internal/service/albumService"

	"taylor-music-back/internal/model/entity"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Controller 专辑控制器
type ControllerV1 struct{}

// 创建一个控制器实例
var Album = &ControllerV1{}

func NewAlbumV1() api.AlbumApiV1 {
	return &ControllerV1{}
}
func getAlbumService() albumService.AlbumService {
	return albumService.NewAlbumService()
}

// List 获取专辑列表
func (c *ControllerV1) List(ctx context.Context, _ *request.AlbumListReq) (res *request.AlbumListRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	getAlbumService().GetAlbumList(req)
	return res, err
}

// Create 创建专辑
func (c *ControllerV1) Create(ctx context.Context, getData *request.AlbumCreateReq) (res *request.AlbumCreateRes, err error) {
	req := ghttp.RequestFromCtx(ctx)

	if err == nil {
		getAlbumService().CreateAlbum(req, getData)
	}
	return res, err
}

// Edit 更新专辑
func (c *ControllerV1) Edit(ctx context.Context, _ *request.AlbumEditReq) (res *request.AlbumEditRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	albumEditVO := entity.AlbumEditVO{}
	err = req.GetRequestStruct(&albumEditVO)
	if err == nil {
		getAlbumService().UpdateAlbum(req, albumEditVO)
	}
	return res, err
}


// Delete 删除专辑
func (c *ControllerV1) Delete(ctx context.Context, _ *request.AlbumDeleteReq) (res *request.AlbumDeleteRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	albumDeleteVO := entity.AlbumDeleteVO{}

	//解析请求体，解释失败则err不为空
	err = req.GetRequestStruct(&albumDeleteVO)
	if err == nil {
		getAlbumService().DeleteAlbum(req, albumDeleteVO)
	}
	return res, err
}
