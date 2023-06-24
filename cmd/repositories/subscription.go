package repositories

import (
	"context"
	"log"

	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type SubscriptionRepository struct {
	collection *mongo.Collection
}

func NewSubscriptionRepository(database *mongo.Database, collectionName string) *SubscriptionRepository {
	collection := database.Collection(collectionName)
	return &SubscriptionRepository{
		collection: collection,
	}
}

func (r *SubscriptionRepository) Create(subscription *entities.Subscription) error {
	_, err := r.collection.InsertOne(context.TODO(), subscription)
	if err != nil {
		log.Println("Error creating subscription:", err)
		return err
	}
	return nil
}

func (r *SubscriptionRepository) Update(subscription *entities.Subscription) error {
	filter := bson.M{"_id": subscription.ID}
	update := bson.M{"$set": subscription}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating subscription:", err)
		return err
	}
	return nil
}

func (r *SubscriptionRepository) Delete(id string) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting subscription:", err)
		return err
	}
	return nil
}

func (r *SubscriptionRepository) FindByID(id string) (*entities.Subscription, error) {
	filter := bson.M{"_id": id}
	var subscription entities.Subscription
	err := r.collection.FindOne(context.TODO(), filter).Decode(&subscription)
	if err != nil {
		log.Println("Error finding subscription:", err)
		return nil, err
	}
	return &subscription, nil
}

func (r *SubscriptionRepository) FindAll() ([]*entities.Subscription, error) {
	var subscriptions []*entities.Subscription
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error finding subscriptions:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var subscription entities.Subscription
		err := cursor.Decode(&subscription)
		if err != nil {
			log.Println("Error decoding subscription:", err)
			return nil, err
		}
		subscriptions = append(subscriptions, &subscription)
	}
	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over subscriptions:", err)
		return nil, err
	}
	return subscriptions, nil
}
