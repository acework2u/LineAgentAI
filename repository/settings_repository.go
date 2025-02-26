package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
	log.Println(id)
	// get result app setting
	appSetting := AppSettings{}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		return res.Err()
	}
	err = res.Decode(&appSetting)

	log.Println("Repo MemberType")
	log.Println(memberType)
	appMemberType := []*MemberTypeSettingImpl{}
	for _, member := range appSetting.MemberType {
		//if member.Title != memberType.Title {
		//	appMemberType = append(appMemberType, member)
		//}
		if member.Title == memberType.Title {
			return nil
		}
		appMemberType = append(appMemberType, member)
	}

	appMemberType = append(appMemberType, memberType)
	//log.Println(appMemberType)

	if err != nil {
		return err
	}
	//res = r.appSettingsCollection.FindOneAndUpdate(r.ctx, bson.D{{"_id", id}}, bson.M{"$push": bson.D{{"memberTypes", memberType}}})
	res = r.appSettingsCollection.FindOneAndUpdate(r.ctx, bson.D{{"_id", id}}, bson.M{"$set": bson.D{{"members_type", appMemberType}}})
	if res.Err() != nil {
		return res.Err()
	}
	// set Title is Indexing
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"title", 1}},
	}
	_, err = r.appSettingsCollection.Indexes().CreateOne(r.ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
func (r *settingsRepository) UpdateMemberType(appId string, memberType *MemberTypeSettingImpl) error {
	appSetting := AppSettings{}
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return err
	}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		return res.Err()
	}
	err = res.Decode(&appSetting)
	if err != nil {
		return err
	}
	// update member type in app setting
	appMemberType := []*MemberTypeSettingImpl{}
	for _, member := range appSetting.MemberType {
		if member.Title == memberType.Title {
			member.Title = memberType.Title
			member.Status = memberType.Status
		}
		appMemberType = append(appMemberType, member)
	}
	// update to database
	res = r.appSettingsCollection.FindOneAndUpdate(r.ctx, bson.D{{"_id", id}}, bson.M{"$set": bson.D{{"members_type", appMemberType}}})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
func (r *settingsRepository) DeleteMemberType(appId string, memberType *MemberTypeSettingImpl) error {
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return err
	}
	res := r.appSettingsCollection.FindOneAndUpdate(r.ctx, bson.D{{"_id", id}}, bson.M{"$pull": bson.D{{"members_type", memberType}}})
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}
func (r *settingsRepository) MemberTypesetting(appId string) ([]*MemberTypeSettingImpl, error) {
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return nil, err
	}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var appSetting AppSettings
	err = res.Decode(&appSetting)
	if err != nil {
		return nil, err
	}
	return appSetting.MemberType, nil
}
func (r *settingsRepository) DeleteAppSettings(appId string) error {
	id, _ := primitive.ObjectIDFromHex(appId)
	err := r.appSettingsCollection.FindOneAndDelete(r.ctx, bson.D{{"_id", id}})
	if err != nil {
		return err.Err()
	}

	return nil
}
func (r *settingsRepository) AddClinicSetting(appId string, clinicSetting *ClinicSettingImpl) error {
	return nil
}
func (r *settingsRepository) AddCourse(appId string, course *Course) error {
	return nil
}
func (r *settingsRepository) AddCourseType(appId string, courseType string) error {
	return nil
}
