package handlers

import "go.mongodb.org/mongo-driver/bson"

func ParamsFilter(q *Queries) bson.M {
	filter := bson.M{}

	if q.Client != "" {
		filter["client"] = q.Client
	}
	if q.Method != "" {
		filter["paymentMethod"] = q.Method
	}
	if q.Car != "" {
		filter["carNumber"] = q.Car
	}
	if q.Date != "" {
		filter["date"] = q.Date
	}
	if q.DeliveryDate != "" {
		filter["deliveryDate"] = q.DeliveryDate
	}

	return filter
}
