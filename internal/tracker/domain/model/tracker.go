package model

type StudySession struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userId"`
	Name         string `json:"name"`
	SessionStart string `json:"sessionStart"`
	SessionEnd   string `json:"sessionEnd"`
	TotalTime    int    `json:"totalTime"`
	Note         string `json:"note"`
}

type StudySessionRequest struct {
	UserId int    `json:"userId" validate:"required"`
	Name   string `json:"name"`
	Note   string `json:"note"`
}
