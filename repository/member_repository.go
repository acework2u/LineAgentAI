package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type memberRepositoryImpl struct {
	ctx              context.Context
	memberCollection *mongo.Collection
}

func NewMemberRepository(ctx context.Context, memberCollection *mongo.Collection) MemberRepository {
	return &memberRepositoryImpl{
		ctx:              ctx,
		memberCollection: memberCollection,
	}
}
func (r *memberRepositoryImpl) CreateMember(member *Member) error {
	memberInfo := Member{}

	err := r.memberCollection.FindOne(r.ctx, bson.M{"lineid": member.LineId}).Decode(&memberInfo)
	if err == nil {
		return errors.New("member already exists")
	}
	filter := bson.D{{"$regex", member.Name}, {"$options", "i"}}
	query := bson.D{{"$or", bson.A{
		bson.D{{"name", filter}},
		bson.D{{"lastname", filter}},
	}}}
	err = r.memberCollection.FindOne(r.ctx, query).Decode(&memberInfo)
	if err == nil {
		return errors.New("user already exists")
	}
	// Create a new user
	result, err := r.memberCollection.InsertOne(r.ctx, member)
	if err != nil {
		var er mongo.WriteException
		if errors.As(err, &er) && er.WriteErrors[0].Code == 11000 {
			return errors.New("member already exists")
		}
		return err
	}

	if result.InsertedID == nil {
		return errors.New("member not created")
	}
	// Set Index
	opt := options.Index().SetUnique(true)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"lineid", 1}, {"name", 1}, {"lastname", 1}, {"email", 1}},
		Options: opt,
	}
	_, err = r.memberCollection.Indexes().CreateOne(r.ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}
func (r *memberRepositoryImpl) GetMemberByLineId(lineId string) (*Member, error) {

	memberRes := r.memberCollection.FindOne(r.ctx, bson.M{"lineid": lineId})
	if memberRes.Err() != nil {
		return nil, memberRes.Err()
	}
	var member Member
	err := memberRes.Decode(&member)
	if err != nil {
		return nil, err
	}
	return &member, nil
}
func (r *memberRepositoryImpl) UpdateMember(lineId string, member *Member) error {

	sr := r.memberCollection.FindOneAndUpdate(r.ctx, bson.M{"lineid": member.LineId}, bson.M{"$set": member})
	if sr.Err() != nil {
		return sr.Err()
	}
	return nil
}
func (r *memberRepositoryImpl) DeleteMember(id string) error {
	return nil
}
func (r *memberRepositoryImpl) CreateJoinEvent(event *JoinEventImpl) error {

	panic("implement me")
}
func (r *memberRepositoryImpl) GetJoinEvent(eventId string) (*JoinEventImpl, error) {
	panic("implement me")
}
func (r *memberRepositoryImpl) CheckJoinEvent(eventId string, userId string) (bool, error) {
	panic("implement me")
}
func (r *memberRepositoryImpl) MemberList() ([]*Member, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"updatedDate", -1}})
	cursor, err := r.memberCollection.Find(r.ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)
	members := []*Member{}
	for cursor.Next(r.ctx) {
		var member Member
		err := cursor.Decode(&member)
		if err != nil {
			return nil, err
		}
		members = append(members, &member)
	}

	return members, nil

}
func (r *memberRepositoryImpl) GetMembers(filter Filter) ([]*Member, error) {
	if filter.Limit <= 0 {
		filter.Limit = 100
	}
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Sort == "" {
		filter.Sort = "desc"
	}
	limit := int64(filter.Limit)
	skip := int64((filter.Page - 1) * filter.Limit)
	sort := int64(1)
	if filter.Sort == "desc" {
		sort = -1
	}
	query := bson.D{}
	if filter.Keyword != "" {
		keyFilter := bson.D{{"$regex", filter.Keyword}, {"$options", "i"}}
		query = append(query, bson.E{Key: "$or", Value: bson.A{"", bson.D{{"name", keyFilter}}, bson.D{{"lastname", keyFilter}}, bson.D{{"email", keyFilter}}}})
	}
	opts := options.Find().SetSort(bson.D{{"updatedDate", sort}}).SetLimit(limit).SetSkip(skip)
	cursor, err := r.memberCollection.Find(r.ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)
	members := []*Member{}
	for cursor.Next(r.ctx) {
		var member Member
		err := cursor.Decode(&member)
		if err != nil {
			return nil, err
		}
	}

	return members, nil

}
func (r *memberRepositoryImpl) UpdateMemberStatus(id string, status bool) error {

	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := r.memberCollection.UpdateOne(r.ctx, bson.M{"_id": uid}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("member not found")
	}

	return nil

}
