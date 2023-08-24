package providers

import "github.com/fedeveron01/golang-base/cmd/core/entities"

type SuscriptionProvider interface {
	CreateSubscription(subscription entities.Subscription) (entities.Subscription, error)
}
