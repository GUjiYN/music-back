package album

import (
	"context"
	v1 "taylor-music-back/api/album/v1"
	"taylor-music-back/internal/service"
)

func (c *ControllerV1) GetAlbumList(ctx context.Context, req *v1.GetAlbumListReq) (res *v1.GetAlbumListRes, err error) {
	// 调用逻辑层方法获取专辑列表
	albums, err := service.Album().GetAlbumList(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	// 构建响应
	res = &v1.GetAlbumListRes{
		Total: len(albums),
		Page:  req.Page,
		Size:  req.Size,
		Pages: (len(albums) + req.Size - 1) / req.Size, // 计算总页数
		List:  albums,                                  // 更新响应中的 List 字段
	}

	return res, nil
}
