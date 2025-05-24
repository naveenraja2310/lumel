package repo

import (
	"context"
	"lumel/internal/database"
	"lumel/internal/model"
	"lumel/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func BulkInsertCustomers(customers []model.Customer) {
	if len(customers) == 0 {
		return
	}

	var models []mongo.WriteModel
	for _, c := range customers {
		filter := bson.M{"customer_id": c.CustomerID}
		update := bson.M{"$set": c}
		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		models = append(models, model)
	}

	result, err := database.CustomerCollection.BulkWrite(context.Background(), models, options.BulkWrite().SetOrdered(false))
	if err != nil {
		logger.Log.Info("Error in insert product", zap.Error(err))
	}

	logger.Log.Info("insert customer summary",
		zap.Int64("InsertedCount", result.InsertedCount),
		zap.Int64("MatchedCount", result.MatchedCount),
		zap.Int64("ModifiedCount", result.ModifiedCount),
		zap.Int64("DeletedCount", result.DeletedCount),
		zap.Int64("UpsertedCount", result.UpsertedCount),
		zap.Any("UpsertedIDs", result.UpsertedIDs),
	)
}
