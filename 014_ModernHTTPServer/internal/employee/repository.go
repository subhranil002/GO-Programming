package course

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Employee, error)
	FindByID(ctx context.Context, id bson.ObjectID) (Employee, error)
	Insert(ctx context.Context, employee Employee) (Employee, error)
	UpdateByID(ctx context.Context, id bson.ObjectID, update bson.D) (Employee, error)
	DeleteByID(ctx context.Context, id bson.ObjectID) (int64, error)
}

type mongoRepository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &mongoRepository{collection: collection}
}

func (r *mongoRepository) FindAll(ctx context.Context) ([]Employee, error) {
	cursor, err := r.collection.Find(
		ctx,
		bson.D{},
		options.Find().SetSort(bson.D{{Key: "name", Value: 1}}),
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	employees := make([]Employee, 0)

	if err := cursor.All(ctx, &employees); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id bson.ObjectID) (Employee, error) {
	var employee Employee

	err := r.collection.FindOne(
		ctx,
		bson.D{{Key: "_id", Value: id}},
	).Decode(&employee)

	return employee, err
}

func (r *mongoRepository) Insert(ctx context.Context, employee Employee) (Employee, error) {
	_, err := r.collection.InsertOne(ctx, employee)
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id bson.ObjectID, update bson.D) (Employee, error) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated Employee

	err := r.collection.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "_id", Value: id}},
		bson.D{{Key: "$set", Value: update}},
		opts,
	).Decode(&updated)

	return updated, err
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id bson.ObjectID) (int64, error) {
	result, err := r.collection.DeleteOne(
		ctx,
		bson.D{{Key: "_id", Value: id}},
	)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
