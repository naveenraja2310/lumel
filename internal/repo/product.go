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

func BulkInsertProducts(products []model.Product) {
	if len(products) == 0 {
		return
	}

	var models []mongo.WriteModel
	for _, p := range products {
		filter := bson.M{"product_id": p.ProductID}
		update := bson.M{"$set": p}
		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		models = append(models, model)
	}

	result, err := database.ProductCollection.BulkWrite(context.Background(), models, options.BulkWrite().SetOrdered(false))
	if err != nil {
		logger.Log.Info("Error in insert product", zap.Error(err))
	}

	logger.Log.Info("insert product summary",
		zap.Int64("InsertedCount", result.InsertedCount),
		zap.Int64("MatchedCount", result.MatchedCount),
		zap.Int64("ModifiedCount", result.ModifiedCount),
		zap.Int64("DeletedCount", result.DeletedCount),
		zap.Int64("UpsertedCount", result.UpsertedCount),
		zap.Any("UpsertedIDs", result.UpsertedIDs),
	)
}
