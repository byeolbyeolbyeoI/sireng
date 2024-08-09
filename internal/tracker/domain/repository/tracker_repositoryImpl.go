package repository

import (
	"database/sql"
	"errors"

	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
	"github.com/chaaaeeee/sireng/util"
)

type trackerRepositoryImpl struct {
	db   *sql.DB
	util util.Util
}

var ()

func NewTrackerRepository(db *sql.DB, util util.Util) TrackerRepository {
	return &trackerRepositoryImpl{
		db:   db,
		util: util,
	}
}

func (t *trackerRepositoryImpl) GetStudySessionsByUserId(userId int) ([]trackerModel.StudySession, error) {
	var studySession trackerModel.StudySession
	var studySessions []trackerModel.StudySession
	rows, err := t.db.Query("SELECT id, user_id, name, session_start, session_end, total_time, note FROM study_sessions WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&studySession.Id,
			&studySession.UserId,
			&studySession.Name,
			&studySession.SessionStart,
			&studySession.SessionEnd,
			&studySession.TotalTime,
			&studySession.Note,
		)
		if err != nil {
			return nil, err
		}

		studySessions = append(studySessions, studySession)
	}

	return studySessions, nil
}

func (t *trackerRepositoryImpl) GetStudySessions() ([]trackerModel.StudySession, error) {
	var studySession trackerModel.StudySession
	var studySessions []trackerModel.StudySession
	rows, err := t.db.Query("SELECT id, user_id, name, session_start, session_end, total_time, note FROM study_sessions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&studySession.Id,
			&studySession.UserId,
			&studySession.Name,
			&studySession.SessionStart,
			&studySession.SessionEnd,
			&studySession.TotalTime,
			&studySession.Note,
		)
		if err != nil {
			return nil, err
		}

		studySessions = append(studySessions, studySession)
	}

	return studySessions, nil
}

func (t *trackerRepositoryImpl) IsSessionActiveByUserId(userId int) (bool, error) {
	var id int
	err := t.db.QueryRow("SELECT id FROM study_sessions WHERE user_id=? AND session_end IS NULL", userId).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (t *trackerRepositoryImpl) CreateStudySession(studySession trackerModel.StudySessionRequest) error {
	_, err := t.db.Exec("INSERT INTO study_sessions(user_id, name, session_start, note) VALUES (?,?,CURRENT_TIMESTAMP, ?)",
		studySession.UserId,
		studySession.Name,
		studySession.Note,
	)
	if err != nil {
		return err
	}

	return nil
}

func (t *trackerRepositoryImpl) EndStudySession(userId int) error {
	_, err := t.db.Exec("UPDATE study_sessions SET session_end=CURRENT_TIMESTAMP WHERE user_id=?", userId)
	if err != nil {
		return err
	}
	return nil
}
