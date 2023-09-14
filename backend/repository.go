package main

import (
	"context"
	"log"
	"monthly-income-expense/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (repository *Repository) GetSalaries() ([]models.Salary, error) {
	collection := repository.client.Database("salary").Collection("salary")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	salaries := []models.Salary{}
	for cur.Next(ctx) {
		var salary models.Salary
		err := cur.Decode(&salary) //& koy
		if err != nil {
			log.Fatal(err)
		}
		salaries = append(salaries, salary)
	}

	return salaries, nil
}

func (repository *Repository) GetSalary(ID string) (models.Salary, error) {
	collection := repository.client.Database("salary").Collection("salary")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	salary := models.Salary{}
	err := collection.FindOne(ctx, bson.M{"id": ID}).Decode(&salary)

	if err != nil {
		log.Fatal(err)
	}
	return salary, nil
}

func (repository *Repository) PostSalary(salary models.Salary) error {
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
