package taskdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const taskCollectionName = "tasks"

var taskCollection *mongo.Collection
var ctx context.Context

type Task struct {
	Id primitive.ObjectID `bson:"_id"`
	Value string `bson:"value"`
}

func Init(db *mongo.Database, nCtx context.Context) {
	ctx = nCtx
	taskCollection = db.Collection(taskCollectionName)
}

func CreateTask(task string) error {
	taskStruct := Task{Value: task}
	_, err := taskCollection.InsertOne(ctx, taskStruct)
	return err
}

func AllTasks() ([]Task, error) {
	var tasks []Task

	cursor, err := taskCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTask(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := taskCollection.DeleteOne(ctx, filter)
	return err
}