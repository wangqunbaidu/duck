package eventhandler

import (
	"duck/internal/models"
	"duck/internal/models/constants"
	"duck/internal/pkg/event"
	"duck/internal/services"
	"reflect"
)

func init() {
	event.RegHandler(reflect.TypeOf(event.FollowEvent{}), handleFollowEvent)
}

func handleFollowEvent(i interface{}) {
	e := i.(event.FollowEvent)

	// 将该用户下的帖子添加到信息流
	services.TopicService.ScanByUser(e.OtherId, func(topics []models.Topic) {
		for _, topic := range topics {
			if topic.Status != constants.StatusOk {
				continue
			}
			_ = services.UserFeedService.Create(&models.UserFeed{
				UserId:     e.UserId,
				DataType:   constants.EntityTopic,
				DataId:     topic.Id,
				AuthorId:   topic.UserId,
				CreateTime: topic.CreateTime,
			})
		}
	})
}
