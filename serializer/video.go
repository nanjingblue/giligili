package serializer

import "singo/model"

// Video 视频的序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title		string `json:"title"`
	Info		string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化用户
func BuildVideo (video model.Video) Video {
	return Video{
		ID:        video.ID,
		Title:	video.Title,
		Info: 	video.Info,
		CreatedAt: video.CreatedAt.Unix(),
	}
}