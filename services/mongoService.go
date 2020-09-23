package services

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/_url-Short-App/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	stringConn string
	nameColl   string
	nameDb     string
}

func NewMongoService() *MongoService {
	m := new(MongoService)
	// m.stringConn = "mongodb+srv://owner:mongoagosto2020@cluster0.ux6pq.mongodb.net/<dbname>?retryWrites=true&w=majority"
	m.stringConn = os.Getenv("DBCONN")
	m.nameDb = "urls"
	m.nameColl = "urls"
	return m
}

func (m *MongoService) Add(url entities.Url) error {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(m.nameDb).Collection(m.nameColl)

	insertResult, err := collection.InsertOne(context.TODO(), url)
	_ = insertResult
	return err
}

func (m *MongoService) FindById(id int) (entities.Url, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(m.nameDb).Collection(m.nameColl)

	filter := bson.D{{"id", id}}
	var result entities.Url
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func (m *MongoService) Find(url entities.Url) (entities.Url, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(m.nameDb).Collection(m.nameColl)

	filter := bson.D{{"text", url.Text}}
	var result entities.Url
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func (m *MongoService) Delete(id int) error {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(m.nameDb).Collection(m.nameColl)

	delFilter := bson.D{{"id", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), delFilter)
	_ = deleteResult
	return err
}

func (m *MongoService) GetAll() ([]entities.Url, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	var results []entities.Url
	collection := client.Database(m.nameDb).Collection(m.nameColl)
	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return results, err
	}

	for cur.Next(context.TODO()) {

		var s entities.Url
		err := cur.Decode(&s)
		if err != nil {
			return results, err
		}

		results = append(results, s)
	}

	return results, err
}

func (m *MongoService) Update(int, entities.Url) error {
	return nil
}

func (m *MongoService) FindLast() (entities.Url, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	opts := options.Find()
	opts.SetSort(bson.D{{"id", -1}})
	opts.SetLimit(1)
	collection := client.Database(m.nameDb).Collection(m.nameColl)
	cur, err := collection.Find(context.TODO(), bson.D{{}}, opts)
	var s entities.Url
	cur.Next(context.TODO())
	err2 := cur.Decode(&s)
	return s, err2
}
