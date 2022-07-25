package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"logistics-go/mongodb"
)

const Pagesize int64 = 10

func InsertShipper(ctx *fiber.Ctx) error {
	shipper := new(Shipper)
	_ = ctx.BodyParser(shipper)
	res, _ := mongodb.Shippers.InsertOne(ctx.Context(), shipper)

	return ctx.JSON(res)
}

func GetShippers(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	q := Queries{}
	_ = ctx.QueryParser(&q)

	filter := ParamsFilter(&q)

	opts := options.Find().SetLimit(Pagesize)
	if q.Page > 0 {
		opts.SetSkip(Pagesize * (q.Page - 1))
	}

	cursor, e := mongodb.Shippers.Find(c, filter, opts)
	if e != nil {
		log.Fatal(e)
	}
	count, _ := mongodb.Shippers.CountDocuments(c, filter)

	shippers := make([]Shipper, 0)

	if err := cursor.All(c, &shippers); err != nil {
		log.Fatal(err)
	}
	_ = cursor.Close(c)

	r := make(map[string]interface{}, 0)
	r["shippers"] = shippers
	r["count"] = count

	return ctx.JSON(r)
}

func DeleteShipper(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	filter := bson.M{"_id": id}
	res, _ := mongodb.Shippers.DeleteOne(context.TODO(), filter)

	return ctx.JSON(res)
}

func GetArrears(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	groupStage := bson.D{{
		"$group",
		bson.M{
			"_id":      "$client",
			"money":    bson.M{"$sum": "$money"},
			"received": bson.M{"$sum": "$received"},
		}}}
	cursor, e := mongodb.Shippers.Aggregate(c, mongo.Pipeline{groupStage})
	if e != nil {
		log.Fatal(e)
	}

	arrears := make([]Arrears, 0)
	for cursor.Next(context.TODO()) {
		var elem Arrears
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		arrears = append(arrears, elem)
	}
	err := cursor.Close(c)
	if err != nil {
		log.Fatal(err)
	}

	return ctx.JSON(arrears)
}
