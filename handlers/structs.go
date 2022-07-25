package handlers

type Queries struct {
	Page         int64  `query:"page"`
	Client       string `query:"client"`
	Method       string `query:"method"`
	Car          string `query:"car"`
	Date         string `query:"date"`
	DeliveryDate string `query:"deliveryDate"`
}

type Shipper struct {
	Id            string `json:"id" bson:"_id"`
	Date          string `json:"date"`
	Client        string `json:"client"`
	Product       string `json:"product"`
	Quantity      int    `json:"quantity"`
	Area          int    `json:"area"`
	Money         int    `json:"money"`
	Received      int    `json:"received"`
	PaymentMethod string `json:"paymentMethod" bson:"paymentMethod"`
	DeliveryDate  string `json:"deliveryDate" bson:"deliveryDate"`
	CarNumber     string `json:"carNumber" bson:"carNumber"`
	Driver        string `json:"driver"`
	Follower      string `json:"follower"`
}

type Arrears struct {
	Client   string `json:"client" bson:"_id"`
	Money    int    `json:"money"`
	Received int    `json:"received"`
}
