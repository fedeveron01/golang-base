package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type SessionGatewayImpl struct {
	sessionRepository repositories.SessionRepository
}

func NewSessionGateway(sessionRepository repositories.SessionRepository) *SessionGatewayImpl {
	return &SessionGatewayImpl{
		sessionRepository: sessionRepository,
	}
}

func (i *SessionGatewayImpl) CreateSession(session entities.Session) error {
	sessionDB := gateway_entities.Session{
		UserId: session.User.ID,
		User: gateway_entities.User{
			UserName: session.User.UserName,
			Password: session.User.Password,
		},
	}

	err := i.sessionRepository.CreateSession(sessionDB)
	if err != nil {
		return err
	}
	return nil
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

func (i *SessionGatewayImpl) DeleteSession(id string) error {
	return i.sessionRepository.DeleteSession(id)
}
