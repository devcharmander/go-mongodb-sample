package database

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Habit struct {
	ID        primitive.ObjectID   `bson:"_id"`
	Name      string               `bson:"name"`
	Track     map[int32]bool       `bson:"track"`
	Reward    string               `bson:"reward"`
	Startdate *timestamp.Timestamp `bson:"start_date"`
}

type MongoClient struct {
	Actions
	client     *mongo.Client
	Collection *mongo.Collection
}

type MongoRequest struct {
	Filter interface{}
	Data   []*Habit
}

var Client *MongoClient

func init() {
	Client = NewClient()
	Client.AddCollection("habitTracker", "habits")
}

func setContext() (context.Context, func()) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

//NewClient returns a db client
func NewClient() *MongoClient { //TODO client struct needs to be decided
	client, err := mongo.NewClient(options.Client().ApplyURI(""))
	if err != nil {
		log.Fatalf("Error - unable to create a new mongo client.Error: %v", err)
		//return nil, err
	}
	return &MongoClient{
		client: client,
	}

}

//AddCollections Creates collections at specified database
func (c *MongoClient) AddCollection(database, name string) {

	db := c.client.Database(database)

	collection := db.Collection(name)

	Client.Collection = collection

}

func (c *MongoClient) Create(data interface{}) ([]*Habit, error) {
	var err error
	context, cancel := setContext()
	defer cancel()
	req := c.getRequestObj(data)
	c.client.Connect(context)
	defer c.client.Disconnect(context)
	for _, v := range req.Data {
		v.ID = primitive.NewObjectID()
		_, err = Client.Collection.InsertOne(context, v)
		if err != nil {
			log.Println("Could not insert record", v)
		}
	}
	return nil, err
}

func (c *MongoClient) Retrieve(data interface{}) ([]*Habit, error) {
	var cur *mongo.Cursor
	var habits []*Habit
	var err error
	context, cancel := setContext()
	defer cancel()
	req := c.getRequestObj(data)

	c.client.Connect(context)
	defer c.client.Disconnect(context)

	if req.Filter == nil {
		cur, err = c.Collection.Find(context, bson.D{{}})
		if err != nil {
			log.Fatal("Could not get the collection")
		}
	} else {
		cur, err = c.Collection.Find(context, req.Filter)
		if err != nil {
			log.Fatal("Could not get the collection")
		}
	}

	for cur.Next(context) {
		var h Habit
		err = cur.Decode(&h)
		if err != nil {
			log.Fatalf("Could not get the Habit. Error %v", err)
		}
		habits = append(habits, &h)
	}

	return habits, err

}

// func (c *MongoClient) Update(data interface{}) (interface{}, error) {

// }
// func (c *MongoClient) Delete(data interface{}) (interface{}, error) {

// }

func (c *MongoClient) getRequestObj(data interface{}) *MongoRequest {
	log.Printf("data:  %v", data)
	req, ok := data.(*MongoRequest)
	if !ok {
		log.Fatalf("Error: Not a valid Mongo Request %v %v", ok, req)
	}
	return req
}
