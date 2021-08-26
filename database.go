package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBConn *mongo.Client
	ChallengesCol *mongo.Collection
	ChannelsCol *mongo.Collection
	DataCol *mongo.Collection
)

const DBName = "issueconsumer"

type Challenge struct {
	Id primitive.ObjectID  `bson:"_id"`
	Name string `bson:"name"`
	Github string `bson:"github"`
	Resolved bool `bson:"resolved"`
}

func searchChallenges(q string) *Challenge {
	result := new(Challenge)
	err := ChallengesCol.FindOne(context.TODO(), bson.M{"$text":bson.M{"$search":q}}, &options.FindOneOptions{Sort:bson.M{"score":bson.M{"$meta":"textScore"}}}).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

type Answer struct {
	Answer bool `bson:"answer"`
	Text string `bson:"text"`
	Name string `bson:"name"`
	URL string `bson:"url"`
}

func searchAnswersByChallengeID(id primitive.ObjectID) *Answer {
	answer := new(Answer)
	err := DataCol.FindOne(context.TODO(), bson.M{"challenge":id, "answer":true}).Decode(&answer)
	if err != nil {
		log.Println(err)
	}
	return answer
}

func connectDB() {
	var err error
	DBConn, err = mongo.Connect(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connection successful.")
	db := DBConn.Database(DBName)
	ChallengesCol = db.Collection("challenges")
	ChannelsCol = db.Collection("channels")
	DataCol = db.Collection("data")
}