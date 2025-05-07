package album

import (
	"context"
	v1 "taylor-music-back/api/album/v1"
	"taylor-music-back/internal/service"
)

func (c *ControllerV1) GetAlbumList(ctx context.Context, req *v1.GetAlbumListReq) (res *v1.GetAlbumListRes, err error) {
	// 调用逻辑层方法获取专辑列表
	albums, totalCount, err := service.Album().GetAlbumList(ctx, req.Page, req.Size, req.Keyword)
	if err != nil {
		return nil, err
	}

	// 构建响应
	res = &v1.GetAlbumListRes{
		Code:    0,
		Message: "",
	}

	// 填充数据部分
	res.Data.Records = albums
	res.Data.Total = totalCount
	res.Data.Size = req.Size
	res.Data.Current = req.Page

	return res, nil
}
