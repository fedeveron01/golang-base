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

func (s *SessionGatewayImpl) CreateSession(session entities.Session) (entities.Session, error) {
	sessionDB := s.ToServiceEntity(session)

	sessionDB, err := s.sessionRepository.CreateSession(sessionDB)
	if err != nil {
		return entities.Session{}, err
	}

	session = s.ToBusinessEntity(sessionDB)
	return session, nil
}

func (s *SessionGatewayImpl) FindAll() ([]entities.Session, error) {
	sessionsDB, err := s.sessionRepository.FindAll()
	if err != nil {
		return nil, err
	}

	sessions := make([]entities.Session, len(sessionsDB))
	for index, sessionDB := range sessionsDB {
		sessions[index] = s.ToBusinessEntity(sessionDB)
	}
	return sessions, nil
}

func (s *SessionGatewayImpl) FindSessionById(id float64) (entities.Session, error) {
	sessionDB, err := s.sessionRepository.FindSessionById(id)
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

func (s *SessionGatewayImpl) SessionIsExpired(id float64) bool {
	return s.sessionRepository.SessionIsExpired(id)

}

func (s *SessionGatewayImpl) UpdateSession(session entities.Session) error {
	sessionDB := gateway_entities.Session{
		UserId: session.User.ID,
		User: gateway_entities.User{
			UserName: session.User.UserName,
			Password: session.User.Password,
		},
	}

	return s.sessionRepository.UpdateSession(sessionDB)
}

func (s *SessionGatewayImpl) DeleteSession(id float64) error {
	return s.sessionRepository.DeleteSession(id)
}

func (s *SessionGatewayImpl) ToBusinessEntity(sessionDB gateway_entities.Session) entities.Session {
	session := entities.Session{
		EntitiesBase: core.EntitiesBase{
			ID: sessionDB.ID,
		},
		User: entities.User{
			UserName: sessionDB.User.UserName,
			Password: sessionDB.User.Password,
		},
	}
	return session
}

func (s *SessionGatewayImpl) ToServiceEntity(session entities.Session) gateway_entities.Session {
	sessionDB := gateway_entities.Session{
		UserId: session.User.ID,
		User: gateway_entities.User{
			UserName: session.User.UserName,
			Password: session.User.Password,
		},
	}
	return sessionDB
}
