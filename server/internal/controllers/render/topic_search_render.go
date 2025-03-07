package render

import (
	"duck/internal/models"
	"duck/internal/pkg/search"
	"duck/internal/services"
)

func BuildSearchTopics(docs []search.TopicDocument) []models.SearchTopicResponse {
	var items []models.SearchTopicResponse
	for _, doc := range docs {
		items = append(items, BuildSearchTopic(doc))
	}
	return items
}

func BuildSearchTopic(doc search.TopicDocument) models.SearchTopicResponse {
	rsp := models.SearchTopicResponse{
		Id:         doc.Id,
		Title:      doc.Title,
		Summary:    doc.Content,
		CreateTime: doc.CreateTime,
		User:       BuildUserInfoDefaultIfNull(doc.UserId),
	}

	if doc.NodeId > 0 {
		node := services.TopicNodeService.Get(doc.NodeId)
		rsp.Node = BuildNode(node)
	}

	tags := services.TopicService.GetTopicTags(doc.Id)
	rsp.Tags = BuildTags(tags)
	return rsp
}
