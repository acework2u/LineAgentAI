package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
