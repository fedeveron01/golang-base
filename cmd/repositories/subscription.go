package repositories

import (
	"gorm.io/gorm"
	"log"

	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(database *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: database,
	}
}

func (r *SubscriptionRepository) Create(subscription *entities.Subscription) error {
	res := r.db.Create(subscription)
	if res.Error != nil {
		log.Println("Error creating subscription with id:", res)
		return res.Error
	}
	return nil
}

func (r *SubscriptionRepository) FindAll() ([]entities.Subscription, error) {
	var subscriptions []entities.Subscription
	res := r.db.Find(&subscriptions)
	if res.Error != nil {
		log.Println("Error finding subscriptions:", res)
		return nil, res.Error
	}
	return subscriptions, nil
}
func (r *SubscriptionRepository) Update(subscription *entities.Subscription) error {
	res := r.db.Save(subscription)
	if res.Error != nil {
		log.Println("Error updating subscription with id:", res)
		return res.Error
	}
	return nil
}

func (r *SubscriptionRepository) Delete(id string) error {
	res := r.db.Delete(&entities.Subscription{}, id)
	if res.Error != nil {
		log.Println("Error deleting subscription with id:", id)
		return res.Error
	}
	return nil
}
