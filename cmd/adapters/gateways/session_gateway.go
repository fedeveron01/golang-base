package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type SessionGateway interface {
	CreateSession(session entities.Session) (entities.Session, error)
	FindAll() ([]entities.Session, error)
	FindSessionById(id float64) (entities.Session, error)
	UpdateSession(session entities.Session) error
	SessionIsExpired(id float64) bool
	DeleteSession(id float64) error
}

type SessionGatewayImpl struct {
	sessionRepository repositories.SessionRepository
}

func NewSessionGateway(sessionRepository repositories.SessionRepository) *SessionGatewayImpl {
	return &SessionGatewayImpl{
		sessionRepository: sessionRepository,
	}
}

func (i *SessionGatewayImpl) CreateSession(session entities.Session) (entities.Session, error) {
	sessionDB := gateway_entities.Session{
		UserId: session.User.ID,
	}

	sessionDB, err := i.sessionRepository.CreateSession(sessionDB)
	if err != nil {
		return entities.Session{}, err
	}

	session = entities.Session{
		EntitiesBase: core.EntitiesBase{
			ID: sessionDB.ID,
		},
		User: entities.User{
			UserName: sessionDB.User.UserName,
			Password: sessionDB.User.Password,
		},
	}
	return session, nil
}

func (i *SessionGatewayImpl) FindAll() ([]entities.Session, error) {
	sessionsDB, err := i.sessionRepository.FindAll()
	if err != nil {
		return nil, err
	}

	sessions := make([]entities.Session, len(sessionsDB))
	for i, sessionDB := range sessionsDB {
		sessions[i] = entities.Session{
			EntitiesBase: core.EntitiesBase{
				ID: sessionDB.ID,
			},
			User: entities.User{
				UserName: sessionDB.User.UserName,
				Password: sessionDB.User.Password,
			},
		}
	}
	return sessions, nil
}

func (i *SessionGatewayImpl) FindSessionById(id float64) (entities.Session, error) {
	sessionDB, err := i.sessionRepository.FindSessionById(id)
	if err != nil {
		return entities.Session{}, err
	}

	session := entities.Session{
		EntitiesBase: core.EntitiesBase{
			ID: sessionDB.ID,
		},
		User: entities.User{
			UserName: sessionDB.User.UserName,
			Password: sessionDB.User.Password,
		},
	}
	return session, nil
}

func (i *SessionGatewayImpl) SessionIsExpired(id float64) bool {
	return i.sessionRepository.SessionIsExpired(id)

}

func (i *SessionGatewayImpl) UpdateSession(session entities.Session) error {
	sessionDB := gateway_entities.Session{
		UserId: session.User.ID,
		User: gateway_entities.User{
			UserName: session.User.UserName,
			Password: session.User.Password,
		},
	}

	return i.sessionRepository.UpdateSession(sessionDB)
}

func (i *SessionGatewayImpl) DeleteSession(id float64) error {
	return i.sessionRepository.DeleteSession(id)
}
