package Controller

import (
	"context"
	"log"
	"time"

	"github.com/kriangkrai/Mee/MongoDB/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadDoc(device string) []Models.DataModel {

	ctx, cancelFindOne := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{"device": device}
	SingleResult, errFind := collection.Find(ctx, filter)
	if errFind != nil {
		panic(errFind)
	}
	cancelFindOne()

	datas := []Models.DataModel{}
	defer SingleResult.Close(ctx)
	for SingleResult.Next(ctx) {
		var episode Models.DataModel
		if err := SingleResult.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		datas = append(datas, episode)
	}

	return datas
}

func InsertDoc(data Models.DataModel) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, Models.DataModel{ID: primitive.NewObjectID(), Device: data.Device, Date: time.Now().Local().String(), Value: data.Value})

	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteDoc(device string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := collection.DeleteMany(ctx, bson.M{"device": device})

	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateDoc(data Models.DataModel) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "value", Value: data.Value},
	}}}
	res, err := collection.UpdateMany(ctx, bson.M{"device": data.Device}, update)

	if err != nil {
		return nil, err
	}
	return res, nil
}
