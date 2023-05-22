package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
)

type SessionRepository interface {
	Insert(newSession *entity.Session) (*entity.Session, error)
	FindByID(id int64) (*entity.Session, error)
	FindAll() ([]entity.Session, error)
	Update(updateSession *entity.Session) (*entity.Session, error)
	Delete(deletedSession *entity.Session) error
}

type sessionRepository struct {
	db *sql.DB
}

func newSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepository{db}
}

// ======================= INSERT ==============================
func (r *sessionRepository) Insert(newSession *entity.Session) (*entity.Session, error) {
	stmt, err := r.db.Prepare("INSERT INTO session(username, token, expired_at, created_at) VALUES ($1,$2,$3,$4) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newSession.Username, newSession.Token, newSession.Expired_At, newSession.Created_At).Scan(&newSession.ID)
	if err != nil {
		return nil, err
	}
	return newSession, nil
}

// ======================= FIND BY ID ==============================
func (r *sessionRepository) FindByID(id int64) (*entity.Session, error) {
	var session entity.Session

	stmt, err := r.db.Prepare("SELECT id, username, token, expired_at, created_at FROM session WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(id).Scan(&session.ID, &session.Username, &session.Token, &session.Expired_At, &session.Created_At)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

// ======================= FIND ALL ==============================
func (r *sessionRepository) FindAll() ([]entity.Session, error) {
	var sessions []entity.Session
	rows, err := r.db.Query("SELECT id, username, token, expired_at, created_at FROM session")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var session entity.Session
		err := rows.Scan(&session.ID, &session.Username, &session.Token, &session.Expired_At, &session.Created_At)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

// ======================= UPDATE ==============================
func (r *sessionRepository) Update(updateSession *entity.Session) (*entity.Session, error) {
	stmt, err := r.db.Prepare("UPDATE session SET username = $1, token = $2, expired_at = $3, created_at = $4 WHERE id = $5")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateSession.Username, updateSession.Token, updateSession.Expired_At, updateSession.Created_At, &updateSession.ID)
	if err != nil {
		return nil, err
	}

	return updateSession, err
}

// ======================= DELETE ==============================
func (r *sessionRepository) Delete(deletedSession *entity.Session) error {
	stmt, err := r.db.Prepare("DELETE FROM session WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deletedSession.ID)
	if err != nil {
		return err
	}

	return nil
}
