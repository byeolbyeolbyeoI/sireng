package repository

import (
	"database/sql"
)

type TrackerRepositoryImpl struct {
	db *sql.DB
}

var ()

func NewTrackerRepository(db *sql.DB) TrackerRepository {
	return &TrackerRepositoryImpl{db: db}
}

func (t *TrackerRepositoryImpl) GetStudyTime() int {

	return -1
}
