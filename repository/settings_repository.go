package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// get result app setting
	appSetting := AppSettings{}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		return res.Err()
	}
	err = res.Decode(&appSetting)

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
		if member.Id == memberType.Id {
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
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return fmt.Errorf("invalid appId: %v", err)
	}
	if course.Id == "" {
		course.Id = primitive.NewObjectID().Hex()
	}
	// Add course to Courses array if it doesn't already exist
	update := bson.M{
		"$addToSet": bson.M{"courses": course},
	}
	res, err := r.appSettingsCollection.UpdateOne(r.ctx, bson.M{"_id": id}, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %v", err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with id: %v", appId)
	}
	return nil
}
func (r *settingsRepository) UpdateCourse(appId string, course *Course) error {
	if course == nil {
		return fmt.Errorf("course cannot be nil")
	}

	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return fmt.Errorf("invalid appId: %v", err)
	}

	appSetting := AppSettings{}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return fmt.Errorf("no app settings found with id: %s", appId)
		}
		return res.Err()
	}

	err = res.Decode(&appSetting)
	if err != nil {
		return fmt.Errorf("failed to decode app settings: %v", err)
	}

	courses := []*Course{}
	found := false
	for _, item := range appSetting.Courses {
		if item.Id == course.Id {
			item.Name = course.Name
			item.Desc = course.Desc
			item.Type = course.Type
			item.Img = course.Img
			item.Status = course.Status
			found = true
		}
		courses = append(courses, item)
	}

	if !found {
		return fmt.Errorf("course with id %s not found", course.Id)
	}

	updateResult, err := r.appSettingsCollection.UpdateOne(r.ctx, bson.D{{"_id", id}}, bson.M{"$set": bson.D{{"courses", courses}}})
	if err != nil {
		return err
	}

	if updateResult.MatchedCount == 0 {
		return fmt.Errorf("no matching document found with id: %s", appId)
	}

	return nil
}
func (r *settingsRepository) DeleteCourse(appId string, course *Course) error {
	if course == nil {
		return fmt.Errorf("course cannot be nil")
	}

	// Convert appId to ObjectID
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return fmt.Errorf("invalid appId: %s, error: %v", appId, err)
	}

	// Use MongoDB's $pull operator to atomically remove the course
	update := bson.M{"$pull": bson.M{"courses": bson.M{"id": course.Id}}}
	res, err := r.appSettingsCollection.UpdateOne(r.ctx, bson.M{"_id": id}, update)
	if err != nil {
		return fmt.Errorf("failed to update app settings: %v", err)
	}

	// Check if the document was found
	if res.MatchedCount == 0 {
		return fmt.Errorf("no app settings found with id %s", appId)
	}

	// Check if the course was removed
	if res.ModifiedCount == 0 {
		return fmt.Errorf("course with id %s not removed (not found in courses array)", course.Id)
	}

	return nil
}
func (r *settingsRepository) CourseListSetting(appId string) ([]*Course, error) {

	id, _ := primitive.ObjectIDFromHex(appId)
	appSetting := AppSettings{}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Decode(&appSetting)
	if err != nil {
		return nil, err
	}
	return appSetting.Courses, nil
}
func (r *settingsRepository) AddCourseType(appId string, courseType *CourseType) error {
	if appId == "" {
		return fmt.Errorf("appId cannot be empty")
	}

	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return fmt.Errorf("invalid appId: %s", appId)
	}

	// Validate course type name and prepare ID
	courseType.Id = primitive.NewObjectID().Hex()

	// Use $addToSet to ensure uniqueness and atomic update
	update := bson.M{
		"$addToSet": bson.M{"course_type": courseType},
	}
	opts := options.FindOneAndUpdate().SetUpsert(false).SetReturnDocument(options.After)
	res := r.appSettingsCollection.FindOneAndUpdate(r.ctx, bson.M{"_id": id}, update, opts)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return fmt.Errorf("no app settings found with id: %s", appId)
		}
		return fmt.Errorf("failed to update app settings: %w", res.Err())
	}
	return nil
}
func (r *settingsRepository) CourseTypeList(appId string) ([]*CourseType, error) {
	if appId == "" {
		return nil, fmt.Errorf("appId cannot be empty")
	}
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return nil, err
	}
	appSetting := AppSettings{}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		return nil, res.Err()
	}
	err = res.Decode(&appSetting)
	if err != nil {
		return nil, err
	}
	return appSetting.CourseType, nil

}
func (r *settingsRepository) UpdateCourseType(appId string, courseType *CourseType) error {
	if appId == "" {
		return fmt.Errorf("appId cannot be empty")
	}
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return err
	}
	appSetting := AppSettings{}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return fmt.Errorf("no app settings found with id: %s", appId)
		}
	}
	err = res.Decode(&appSetting)
	if err != nil {
		return err
	}
	found := false
	for _, course := range appSetting.CourseType {
		if course.Id == courseType.Id {
			found = true
			course = courseType
		}
		appSetting.CourseType = append(appSetting.CourseType, course)
	}
	if !found {
		return fmt.Errorf("course type with id %s not found", courseType.Id)
	}
	update := bson.M{
		"$set": bson.D{{"course_type", appSetting.CourseType}},
	}
	updateResult, err := r.appSettingsCollection.UpdateOne(r.ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 0 {
		return fmt.Errorf("no matching document found with id: %s", appId)
	}
	return nil

}
func (r *settingsRepository) DeleteCourseType(appId string, courseType *CourseType) error {
	if appId == "" {
		return fmt.Errorf("appId cannot be empty")
	}
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return err
	}
	appSetting := AppSettings{}
	res := r.appSettingsCollection.FindOne(r.ctx, bson.D{{"_id", id}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return fmt.Errorf("no app settings found with id: %s", appId)
		}
		return res.Err()
	}
	err = res.Decode(&appSetting)
	if err != nil {
		return err
	}
	courseType.Id = ""
	for _, course := range appSetting.CourseType {
		if course.Id == courseType.Id {
			course.Id = ""
		}
		appSetting.CourseType = append(appSetting.CourseType, course)
	}
	update := bson.M{
		"$set": bson.D{{"course_type", appSetting.CourseType}},
	}
	deleteResult, err := r.appSettingsCollection.UpdateOne(r.ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}
	if deleteResult.MatchedCount == 0 {
		return fmt.Errorf("no matching document found with id: %s", appId)
	}
	return nil
}
