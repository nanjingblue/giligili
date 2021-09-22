package service

import (
	"singo/model"
	"singo/serializer"
)

// ListVideoService 管理视频详情的服务
type ListVideoService struct {
}

// List 展示视频详情
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code: 5000,
			Msg: "数据库来接错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
