package repo

import (
	"context"
	"fmt"
	"lumel/internal/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTotalRevenue(startDate, endDate time.Time) (float64, error) {
	matchStage := bson.M{
		"date_of_sale": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	projectStage := bson.M{
		"total": bson.M{
			"$add": []interface{}{
				bson.M{"$multiply": []interface{}{
					"$quantity_sold",
					"$unit_price",
					bson.M{"$subtract": []interface{}{1, "$discount"}},
				}},
				"$shipping_cost",
			},
		},
	}

	groupStage := bson.M{"_id": nil, "totalRevenue": bson.M{"$sum": "$total"}}

	cursor, err := database.OrderCollection.Aggregate(context.Background(), mongo.Pipeline{
		{{Key: "$match", Value: matchStage}},
		{{Key: "$project", Value: projectStage}},
		{{Key: "$group", Value: groupStage}},
	})
	if err != nil {
		return 0.0, err
	}
	defer cursor.Close(context.Background())

	var result []bson.M
	if err := cursor.All(context.Background(), &result); err != nil {
		return 0.0, err
	}

	totalRevenue := 0.0
	if len(result) > 0 {
		totalRevenue = result[0]["totalRevenue"].(float64)
	}

	return totalRevenue, nil
}

func GroupedRevenue(startDate, endDate time.Time, groupField string) (interface{}, error) {

	match := bson.M{"date_of_sale": bson.M{"$gte": startDate, "$lte": endDate}}
	project := bson.M{
		groupField: fmt.Sprintf("$%s", groupField),
		"total": bson.M{
			"$add": []interface{}{
				bson.M{"$multiply": []interface{}{
					"$quantity_sold",
					"$unit_price",
					bson.M{"$subtract": []interface{}{1, "$discount"}},
				}},
				"$shipping_cost",
			},
		},
	}
	group := bson.M{
		"_id":          fmt.Sprintf("$%s", groupField),
		"totalRevenue": bson.M{"$sum": "$total"},
	}

	cursor, err := database.OrderCollection.Aggregate(context.Background(), mongo.Pipeline{
		{{Key: "$match", Value: match}},
		{{Key: "$project", Value: project}},
		{{Key: "$group", Value: group}},
	})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var result []bson.M
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}
