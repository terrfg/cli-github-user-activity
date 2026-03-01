package main

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		ID           json.Number `json:"id"`
		Login        string      `json:"login"`
		DisplayLogin string      `json:"display_login"`
		GravatarId   string      `json:"gravatar_id"`
		Url          string      `json:"url"`
		AvatarUrl    string      `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   json.Number `json:"id"`
		Name string      `json:"name"`
		Url  string      `json:"url"`
	} `json:"repo"`
	Payload struct {
		RepositoryID json.Number `json:"repository_id"`
		PushID       json.Number `json:"push_id"`
		Ref          string      `json:"ref"`
		Head         string      `json:"head"`
		Before       string      `json:"before"`
	} `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}
