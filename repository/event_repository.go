package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type eventRepositoryImpl struct {
	ctx              context.Context
	eventsCollection *mongo.Collection
}

func NewEventRepository(ctx context.Context, eventsCollection *mongo.Collection) EventsRepository {
	return &eventRepositoryImpl{
		ctx:              ctx,
		eventsCollection: eventsCollection,
	}
}

func (r *eventRepositoryImpl) EventJoin(event *MemberEventImpl) error {
	// Check if the event and user already exist
	existingEvent := MemberEventImpl{}
	err := r.eventsCollection.FindOne(r.ctx, bson.M{
		"eventId": event.EventId,
		"userId":  event.UserId,
	}).Decode(&existingEvent)

	if err == nil {
		// If no error, it means the event and user already exist
		return fmt.Errorf("user already joined this event")
	} else if err != mongo.ErrNoDocuments {
		// If the error is not "no documents", return the error
		return fmt.Errorf("failed to check existing event: %w", err)
	}

	// Proceed with further logic (e.g., insert the event) here
	result, err := r.eventsCollection.InsertOne(r.ctx, event)
	if err != nil {
		return err
	}
	//fmt.Println(result.InsertedID)
	if result.InsertedID == nil {
		return errors.New("event not created")
	}
	return nil
}
func (r *eventRepositoryImpl) EventLeave(event *MemberEventImpl) error {
	panic("implement me")
}
func (r *eventRepositoryImpl) GetEvent(eventId string) (*MemberEventImpl, error) {
	panic("implement me")
}
func (r *eventRepositoryImpl) GetEvents(filter Filter) ([]*MemberEventImpl, error) {
	panic("implement me")
}
func (r *eventRepositoryImpl) CheckJoinEvent(eventId string, userId string) (bool, error) {

	// Create a filter to check if the event exists for the given eventId and userId
	eid := fmt.Sprintf("%s", eventId)
	filter := bson.M{
		"eventId": eid,
		"userId":  userId,
	}

	// Check if the event exists
	count, err := r.eventsCollection.CountDocuments(r.ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to check event membership: %w", err)
	}

	// If count is greater than 0, the user has joined the event
	if count > 0 {
		return true, nil
	}
	return false, nil
}
func (r *eventRepositoryImpl) GetEventJoin(eventId string, userId string) (*MemberEventImpl, error) {

	filter := bson.M{
		"eventId": eventId,
		"userId":  userId,
	}
	event := MemberEventImpl{}
	err := r.eventsCollection.FindOne(r.ctx, filter).Decode(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
func (r *eventRepositoryImpl) CheckInEvent(userId string, eventCheckIn *EventCheckIn) (bool, error) {
	filer := bson.M{
		"eventId": eventCheckIn.EventId,
		"lineId":  userId,
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := r.eventsCollection.FindOneAndUpdate(r.ctx, filer, bson.M{
		"$set": bson.M{
			"checkIn":      eventCheckIn.CheckIn,
			"checkOut":     eventCheckIn.CheckOut,
			"checkInTime":  eventCheckIn.CheckInTime,
			"checkOutTime": eventCheckIn.CheckOutTime,
			"checkInPlace": eventCheckIn.CheckInPlace,
		},
	}, opts)

	if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil

}
func (r *eventRepositoryImpl) EventByUserId(userId string) ([]*MemberEventImpl, error) {

	//Find event is not checkin filter by user Id
	query := bson.M{
		"userId": userId,
		"$or": bson.A{
			bson.M{"checkIn": bson.M{"$exists": false}},
			bson.M{"checkIn": false},
		},
	}

	events := []*MemberEventImpl{}
	cursor, err := r.eventsCollection.Find(r.ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)
	for cursor.Next(r.ctx) {
		var event MemberEventImpl
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	count, _ := r.eventsCollection.CountDocuments(r.ctx, query)
	log.Println(count)
	return events, nil

}
func (r *eventRepositoryImpl) CreateEvent(event *Event) error {

	res, err := r.eventsCollection.InsertOne(r.ctx, event)
	if err != nil {
		return err
	}
	if res.InsertedID == nil {
		return errors.New("event not created")
	}
	return nil

}
func (r *eventRepositoryImpl) UpdateEvent(eventId string, event *Event) error {

	// Create a filter to match the event by its ID
	filter := bson.M{"eventId": eventId}
	// Create the update data using MongoDB's $set operator
	update := bson.M{
		"$set": event,
	}
	// Perform the update operation
	result, err := r.eventsCollection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}
	// Check if any document was modified
	if result.ModifiedCount == 0 {
		return errors.New("event not found or no changes made")
	}
	return nil
}
func (r *eventRepositoryImpl) DeleteEvent(eventId string) error {
	del, err := r.eventsCollection.DeleteOne(r.ctx, bson.M{"eventId": eventId})
	if err != nil {
		return err
	}
	if del.DeletedCount == 0 {
		return errors.New("event not found or no changes made")
	}
	return nil
}
func (r *eventRepositoryImpl) EventByEventId(eventId string) (*Event, error) {
	res := r.eventsCollection.FindOne(r.ctx, bson.M{"eventId": eventId})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var event Event
	err := res.Decode(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil

}
func (r *eventRepositoryImpl) EventsList() ([]*Event, error) {
	filter := bson.M{}
	events := []*Event{}
	cursor, err := r.eventsCollection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)
	for cursor.Next(r.ctx) {
		var event Event
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}
