package services

import (
	auth "money-service/src/services/auth"
	expense "money-service/src/services/expense"
	tag "money-service/src/services/tag"
	user "money-service/src/services/user"
)

type LineService interface {
	CommandManager(payload *LineMessage)
}
type lineService struct {
	userService      user.UserService
	authService      auth.AuthService
	expenseService   expense.ExpenseService
	customTagService tag.TagService
}

type LineMessage struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Timestamp  int64  `json:"timestamp"`
		Source     struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		Message struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}
type ReplyMessage struct {
	ReplyToken string `json:"replyToken"`
	Messages   []Text `json:"messages"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
