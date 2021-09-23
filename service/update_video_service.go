package service

import (
	"singo/model"
	"singo/serializer"
)

// UpdateVideoService 管理视频投稿的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=300"`
}

// Update 创建视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code: 404,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	model.DB.Save(&video)

	return serializer.Response{
		Data: serializer.BuildVideo(video),
		Msg: "视频更新成功",
	}
}
