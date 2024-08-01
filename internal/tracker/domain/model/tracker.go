package model

import (
	"time"
)

type StudySession struct {
	Id           int       `json:"id"`
	UserId       int       `json:"userId"`
	Name         string    `json:"name"`
	SessionStart time.Time `json:"sessionStart"`
	SessionEnd   time.Time `json:"sessionEnd"`
	TotalTime    int       `json:"totalTime"`
	Note         string    `json:"note"`
}

type StudySessionRequest struct {
	UserId int    `json:"userId" validate:"required"`
	Name   string `json:"name"`
	Note   string `json:"note"`
}
