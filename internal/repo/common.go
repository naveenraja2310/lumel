package repo

import "go.mongodb.org/mongo-driver/mongo"

type CommonRepo struct {
	db *mongo.Database
}

func NewCommonRepo(db *mongo.Database) *CommonRepo {
	return &CommonRepo{
		db: db,
	}
}
