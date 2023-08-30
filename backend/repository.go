package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository() *Repository {
	uri := "mongodb+srv://Cluster:bthn998877@cluster0.hnmuy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func GetCleanTestRepository() *Repository {

	repository := NewRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	stockDB := repository.client.Database("Salary")
	stockDB.Drop(ctx)

	return repository
}
