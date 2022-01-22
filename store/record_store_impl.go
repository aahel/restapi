package store

import (
	"context"
	"time"

	"github.com/aahel/restapi/entity"
	"github.com/aahel/restapi/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type RecordStoreImpl struct {
	l  *zap.SugaredLogger
	db *mongo.Database
}

func NewRecordsStore(l *zap.SugaredLogger, db *mongo.Database) *RecordStoreImpl {
	return &RecordStoreImpl{l, db}
}

func (rec *RecordStoreImpl) GetRecords(startDate, endDate time.Time, minCount int64, maxCount int64) ([]*entity.Record, *errors.AppError) {
	records := []*entity.Record{}
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gte": minCount,
					"$lte": maxCount,
				},
			},
		},
	}
	ctx := context.Background()
	cursor, err := rec.db.Collection("records").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &records); err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}
	if len(records) == 0 {
		return nil, errors.RecordNotFound()
	}
	return records, nil
}
