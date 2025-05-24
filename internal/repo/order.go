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

func BulkInsertOrders(orders []model.Order) {
	if len(orders) == 0 {
		return
	}

	var models []mongo.WriteModel
	for _, o := range orders {
		filter := bson.M{"order_id": o.OrderID}
		update := bson.M{"$set": o}
		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		models = append(models, model)
	}

	result, err := database.OrderCollection.BulkWrite(context.Background(), models, options.BulkWrite().SetOrdered(false))
	if err != nil {
		logger.Log.Info("Error in insert order", zap.Error(err))
	}

	logger.Log.Info("insert order summary",
		zap.Int64("InsertedCount", result.InsertedCount),
		zap.Int64("MatchedCount", result.MatchedCount),
		zap.Int64("ModifiedCount", result.ModifiedCount),
		zap.Int64("DeletedCount", result.DeletedCount),
		zap.Int64("UpsertedCount", result.UpsertedCount),
		zap.Any("UpsertedIDs", result.UpsertedIDs),
	)
}
