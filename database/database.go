package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() (*mongo.Database, context.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //will run when  we're finished main

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://saul:1234@cluster0-ooeaq.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	//---------------------//
	client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx) //will run when  we're finished main
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	data1 := client.Database("data")
	//data1 := client.Database("test_db")
	/*
		database, err := client.ListDatabaseNames(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(database)
	*/
	return data1, ctx

}
