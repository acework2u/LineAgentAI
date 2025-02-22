package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type settingsRepository struct {
	ctx                   context.Context
	appSettingsCollection *mongo.Collection
}

func NewSettingsRepository(ctx context.Context, appSettingsCollection *mongo.Collection) AppSettingsRepository {
	return &settingsRepository{
		ctx:                   ctx,
		appSettingsCollection: appSettingsCollection,
	}
}
func (r *settingsRepository) CreateAppSettings(settings *AppSettings) error {
	res, err := r.appSettingsCollection.InsertOne(r.ctx, settings)
	if err != nil {
		return err
	}
	if res.InsertedID == nil {
		return nil
	}
	// create the Indexing for search
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"name", 1}},
	}
	_, err = r.appSettingsCollection.Indexes().CreateOne(r.ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
func (r *settingsRepository) GetAppSettings() (*AppSettings, error) {
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var settings AppSettings
	err := res.Decode(&settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil

}
func (r *settingsRepository) UpdateAppSettings(settings *AppSettings) error {
	appId := settings.Id
	// convert id string to primitive.ObjectId
	res := r.appSettingsCollection.FindOneAndUpdate(r.ctx, bson.D{{"_id", appId}}, bson.M{"$set": settings})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
func (r *settingsRepository) AddMemberType(appId string, memberType *MemberTypeSettingImpl) error {
	// find and update
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return err
	}
	res := r.appSettingsCollection.FindOneAndUpdate(r.ctx, bson.D{{"_id", id}}, bson.M{"$push": bson.D{{"memberTypes", memberType}}})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
