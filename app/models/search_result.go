package models

import (
	"time"
)

type SearchResult struct {
	Title      string             `json:"title" bson:"title"`
	Snippet    string             `json:"snippet" bson:"snippet"`
	Categories []string           `json:"categories" bson:"categories"`
	Attributes map[string][]Value `json:"attributes" bson:"attributes"`
	Links      []string           `json:"links" bson:"links"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	ImageURL   string             `json:"image_url" bson:"image_url"`
}

func (s SearchResult) ResultFrom(entity Entity) SearchResult {

	var dataStructure = make(map[string][]Value)

	for _, attribute := range entity.Attributes {
		dataStructure[attribute.Name] = attribute.Values
	}

	s.Title = entity.Title
	s.Snippet = entity.Snippet
	s.Categories = entity.Categories
	s.Attributes = dataStructure
	s.Links = entity.Links
	s.CreatedAt = entity.CreatedAt
	s.UpdatedAt = entity.UpdatedAt
	s.ImageURL = entity.ImageURL
	return s
}
