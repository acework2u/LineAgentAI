package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type staffRepositoryImpl struct {
	ctx             context.Context
	staffCollection *mongo.Collection
}

func NewStaffRepository(ctx context.Context, staffCollection *mongo.Collection) StaffRepository {
	return &staffRepositoryImpl{
		ctx:             ctx,
		staffCollection: staffCollection,
	}
}
func (r *staffRepositoryImpl) GetStaffs() ([]Staff, error) {

	staffs := []Staff{}
	//opts := options.Find().SetProjection(bson.D{{"_id", 0}})
	cursor, err := r.staffCollection.Find(r.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)
	for cursor.Next(r.ctx) {
		var staff Staff
		err := cursor.Decode(&staff)
		if err != nil {
			return nil, err
		}
		staffs = append(staffs, staff)
	}

	return staffs, nil
}
func (r *staffRepositoryImpl) CreateStaff(staff *Staff) error {

	// insert a new staff
	result, err := r.staffCollection.InsertOne(r.ctx, staff)
	if err != nil {
		return err
	}
	// crate email and name Index
	opt := options.Index().SetUnique(true)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}, {"name", 1}},
		Options: opt,
	}
	_, err = r.staffCollection.Indexes().CreateOne(r.ctx, indexModel)
	if err != nil {
		return err
	}
	if result.InsertedID == nil {
		return errors.New("staff not created")
	}
	return nil
}

func (r *staffRepositoryImpl) UpdateStaff(staff *Staff) error {
	query := bson.D{{"email", staff.Email}}
	res := r.staffCollection.FindOneAndUpdate(r.ctx, query, bson.M{"$set": staff})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
func (r *staffRepositoryImpl) DeleteStaff(staff *Staff) error {
	query := bson.D{{"email", staff.Email}}
	res := r.staffCollection.FindOneAndDelete(r.ctx, query)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
func (r *staffRepositoryImpl) GetStaffById(id string) (*Staff, error) {
	staffId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.D{{"_id", staffId}}
	res := r.staffCollection.FindOne(r.ctx, query)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var staff Staff
	err = res.Decode(&staff)
	if err != nil {
		return nil, err
	}
	return &staff, nil
}
func (r *staffRepositoryImpl) GetStaffByEmail(email string) (*Staff, error) {
	query := bson.D{{"email", email}}
	res := r.staffCollection.FindOne(r.ctx, query)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var staff Staff
	err := res.Decode(&staff)
	if err != nil {
		return nil, err
	}
	return &staff, nil
}
