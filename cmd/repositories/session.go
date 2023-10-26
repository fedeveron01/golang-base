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

func (r *SessionRepository) CreateSession(session gateway_entities.Session) (gateway_entities.Session, error) {
	id := r.db.Create(&session)
	if id.Error != nil {
		return gateway_entities.Session{}, id.Error
	}
	return session, nil
}

func (r *SessionRepository) FindAll() ([]gateway_entities.Session, error) {
	var sessions []gateway_entities.Session
	r.db.Find(&sessions)
	return sessions, nil
}

func (r *SessionRepository) FindSessionById(id float64) (gateway_entities.Session, error) {
	var session gateway_entities.Session
	r.db.Preload("User").First(&session, id)
	return session, nil
}

func (r *SessionRepository) SessionIsExpired(id float64) bool {
	uintId := uint(id)

	var session gateway_entities.Session
	res := r.db.First(&session, uintId)
	if res.Error != nil {
		return true
	}
	return false

}

func (r *SessionRepository) UpdateSession(session gateway_entities.Session) error {
	r.db.Save(&session)
	return nil
}

func (r *SessionRepository) DeleteSession(id float64) error {
	res := r.db.Unscoped().Delete(&gateway_entities.Session{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
