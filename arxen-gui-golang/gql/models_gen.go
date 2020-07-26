// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"time"
)

type Chat struct {
	ChatID         string       `json:"chatId"`
	ClientsIPsList []string     `json:"clientsIPsList"`
	LatestMessage  *TextMessage `json:"latestMessage"`
	ChatAvatar     *string      `json:"chatAvatar"`
	ClientWriting  *string      `json:"clientWriting"`
	ChatName       *string      `json:"chatName"`
}

type Friend struct {
	Nick       *string `json:"nick"`
	UserID     string  `json:"userID"`
	UserIP     *string `json:"userIP"`
	UserAvatar *string `json:"userAvatar"`
	Status     *bool   `json:"status"`
}

type TextMessage struct {
	MessageID string    `json:"messageId"`
	ChatID    string    `json:"chatId"`
	UserNick  *string   `json:"userNick"`
	User      string    `json:"user"`
	TimeStamp time.Time `json:"timeStamp"`
	Text      string    `json:"text"`
}