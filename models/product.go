package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Category    string             `bson:"category"`
	ImageURL    string             `bson:"image_url"`
	Price       float64            `bson:"price"`
	Stock       int                `bson:"stock"`
	Ingredients string             `json:"Ingredients" bson:"ingredients"`
	ShelfLife   int                `bson:"shelf_life"`
	Tags        []string           `bson:"tags"`
	CreatedAt   time.Time          `bson:"created_at"`
}
