package main

import (
	"context"
	"log"
	"monthly-income-expense/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository() *Repository {
	uri := "mongodb+srv://Cluster:bthn998877@cluster0.hnmuy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func (repository *Repository) CreateSalary(salary models.Salary) error {
	collection := repository.client.Database("salary").Collection("salary")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err := collection.InsertOne(ctx, salary)

	if err != nil {
		return err
	}

	return nil
}

func GetCleanTestRepository() *Repository {

	repository := NewRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	stockDB := repository.client.Database("Salary")
	stockDB.Drop(ctx)

	return repository
}
