package repository

import (
	"database/sql"
	"errors"
	"fmt"
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

func (t *trackerRepositoryImpl) GetStudySessionsFromUserId(userId int) ([]trackerModel.StudySession, error) {
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

func (t *trackerRepositoryImpl) CreateSession(studySession trackerModel.StudySessionRequest) error {
	fmt.Println("in")
	fmt.Println(studySession.UserId)
	fmt.Println(studySession.Name)
	fmt.Println(studySession.SessionStart)
	fmt.Println(studySession.Note)
	_, err := t.db.Exec("INSERT INTO study_sessions(user_id, name, session_start, note) VALUES (?,?,?,?)",
		studySession.UserId,
		studySession.Name,
		studySession.SessionStart,
		studySession.Note,
	)
	if err != nil {
		return err
	}

	fmt.Println("bug")
	return nil
}
