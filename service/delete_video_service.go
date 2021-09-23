package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteVideoService 管理视频投稿的服务
type DeleteVideoService struct {
}

// Delete 创建视频
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code: 404,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}
	model.DB.Delete(&video)
	return serializer.Response{
		Data: serializer.BuildVideo(video),
		Msg: "视频删除成功",
	}
}
