package database

import (
	"context"
	"fmt"
	"lumel/internal/model"
	cfg "lumel/pkg/settings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient encapsulates the MongoDB client functionality
type MongoClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var (
	CustomerCollection *mongo.Collection
	ProductCollection  *mongo.Collection
	OrderCollection    *mongo.Collection
)

// NewClient creates a new MongoDB client
func NewClient(config cfg.Configuration) (*MongoClient, error) {
	// Create a context with a 10-second timeout to avoid long-running operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the context is canceled after the operation completes

	// Set client options, including the MongoDB URI and maximum connection pool size
	clientOptions := options.Client().ApplyURI(config.DBURI)

	// Attempt to establish a connection to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err) // Return error if connection fails
	}

	// Get database
	database := client.Database(config.DBName)

	// Initialize collection
	db := client.Database(config.DBName)

	CustomerCollection = db.Collection(model.Customer{}.TableName())
	ProductCollection = db.Collection(model.Product{}.TableName())
	OrderCollection = db.Collection(model.Order{}.TableName())

	return &MongoClient{
		Client:   client,
		Database: database,
	}, nil
}

// Close closes the MongoDB connection
func (c *MongoClient) Close(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}
