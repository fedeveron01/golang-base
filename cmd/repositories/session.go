package repositories

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(database *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: database,
	}
}

func (r *SessionRepository) CreateSession(session entities.Session) error {
	id := r.db.Create(&session)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (r *SessionRepository) FindAll() ([]entities.Session, error) {
	var sessions []entities.Session
	r.db.Find(&sessions)
	return sessions, nil
}

func (r *SessionRepository) UpdateSession(session entities.Session) error {
	r.db.Save(&session)
	return nil
}

func (r *SessionRepository) DeleteSession(id string) error {
	r.db.Delete(&entities.Session{}, id)
	return nil
}
