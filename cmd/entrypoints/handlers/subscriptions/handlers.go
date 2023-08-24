package handler_subscriptions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type CreateSubscriptionHandler struct {
	// use cases
	SubscriptionRepository repositories.SubscriptionRepository
}

type GetAllSubscriptionHandler struct {
	// use cases
	SubscriptionRepository repositories.SubscriptionRepository
}

type GetSubscriptionHandler struct {
	// use cases
	SubscriptionRepository repositories.SubscriptionRepository
}

type EditSubscriptionHandler struct {
	// use cases
	SubscriptionRepository repositories.SubscriptionRepository
}

type DeleteSubscriptionHandler struct {
	// use cases
	SubscriptionRepository repositories.SubscriptionRepository
}

func NewCreateSubscriptionHandler(subscriptionRepository repositories.SubscriptionRepository) entrypoints.Handler {
	return CreateSubscriptionHandler{
		SubscriptionRepository: subscriptionRepository,
	}
}

func NewEditSubscriptionHandler(subscriptionRepository repositories.SubscriptionRepository) entrypoints.Handler {
	return EditSubscriptionHandler{
		SubscriptionRepository: subscriptionRepository,
	}
}

func NewDeleteSubscriptionHandler(subscriptionRepository repositories.SubscriptionRepository) entrypoints.Handler {
	return DeleteSubscriptionHandler{
		SubscriptionRepository: subscriptionRepository,
	}
}

func (p CreateSubscriptionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var subscription entities.Subscription
	json.Unmarshal(reqBody, &subscription)

	json.NewEncoder(w).Encode(subscription)

	fmt.Println(subscription)

	res := p.SubscriptionRepository.Create(&subscription)

	fmt.Fprint(w, res)

}

func (p EditSubscriptionHandler) Handle(w http.ResponseWriter, r *http.Request) {

}

func (p DeleteSubscriptionHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
