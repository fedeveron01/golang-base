package subscription

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type SubscriptionUseCase interface {
	CreateSubscription(subscription entities.Subscription) (entities.Subscription, error)
}

type Implementation struct {
}

func (c Implementation) CreateSubscription(subscription entities.Subscription) (entities.Subscription, error) {
	return entities.Subscription{}, nil
}
