package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
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

func (r *SessionRepository) CreateSession(session gateway_entities.Session) error {
	id := r.db.Create(&session)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (r *SessionRepository) FindAll() ([]gateway_entities.Session, error) {
	var sessions []gateway_entities.Session
	r.db.Find(&sessions)
	return sessions, nil
}

func (r *SessionRepository) UpdateSession(session gateway_entities.Session) error {
	r.db.Save(&session)
	return nil
}

func (r *SessionRepository) DeleteSession(id string) error {
	r.db.Delete(&gateway_entities.Session{}, id)
	return nil
}
