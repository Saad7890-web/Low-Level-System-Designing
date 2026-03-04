package main

import "time"

type Comment struct {
	ID int
	Content string
	Author *User
	CreatedAt time.Time
}

func NewComment(id int, content string, author *User)*Comment{
	return &Comment{
		ID: id,
		Content: content,
		Author: author,
		CreatedAt: time.Now(),
	}
}
