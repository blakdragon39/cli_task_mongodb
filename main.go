package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task/cmd"
	"task/taskdb"
	"time"
)


const taskDb = "taskDb"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer cancel()
	defer disconnectClient(client, ctx)

	if err != nil {
		panic(err)
	}

	db := client.Database(taskDb)

	taskdb.Init(db, ctx)

	cmd.RootCmd.Execute()
}

func disconnectClient(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}