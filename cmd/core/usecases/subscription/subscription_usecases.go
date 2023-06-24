package subscription

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/providers"
)

type SubscriptionUseCase interface {
	CreateSubscription(subscription entities.Subscription) (entities.Subscription, error)
}

type Implementation struct {
	whatsappNotifications providers.WhatsappNotifications
}

func NewUseCase(whatsappNotifications providers.WhatsappNotifications) Implementation {
	return Implementation{
		whatsappNotifications: whatsappNotifications,
	}
}

func (c Implementation) CreateSubscription(subscription entities.Subscription) (entities.Subscription, error) {
	return entities.Subscription{}, nil
}
