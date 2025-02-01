package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
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

	// Filter eventId and userId in Members document of Events
	result := r.eventsCollection.FindOne(r.ctx, bson.M{"eventId": event.EventId})
	if result.Err() != nil {
		return result.Err()
	}
	var eventRes Event
	err := result.Decode(&eventRes)
	if err != nil {
		return err
	}
	members := eventRes.Members
	newMembers := []*MemberEventImpl{}
	for _, member := range members {
		if member.UserId == event.UserId {
			//return errors.New("user already joined this event")
		}
		newMember := MemberEventImpl{
			EventId:        event.EventId,
			UserId:         event.UserId,
			JoinTime:       event.JoinTime,
			Name:           event.Name,
			LastName:       event.LastName,
			Organization:   event.Organization,
			Position:       event.Position,
			Course:         event.Course,
			LineId:         event.LineId,
			LineName:       event.LineName,
			Tel:            event.Tel,
			ReferenceName:  event.ReferenceName,
			ReferencePhone: event.ReferencePhone,
			Clinic:         event.Clinic,
		}
		newMembers = append(newMembers, &newMember)

	}
	newMembers = append(newMembers, event)
	update := bson.M{
		"$set": bson.M{
			"members": newMembers,
		},
	}
	_, err = r.eventsCollection.UpdateOne(r.ctx, bson.M{"eventId": event.EventId}, update)
	if err != nil {
		return err
	}
	return nil

	//
	//// Check Member in Events
	//filter := bson.M{
	//	"members": bson.M{
	//		"$elemMatch": bson.M{
	//			"eventId": event.EventId,
	//			"userId":  event.UserId,
	//		},
	//	},
	//}
	//
	//// Check if the event and user already exist
	//existingEvent := MemberEventImpl{}
	//err := r.eventsCollection.FindOne(r.ctx, filter).Decode(&existingEvent)
	//if err != nil {
	//	return err
	//}
	//
	////err := r.eventsCollection.FindOne(r.ctx, bson.M{
	////	"eventId": event.EventId,
	////	"userId":  event.UserId,
	////}).Decode(&existingEvent)
	//
	//if err == nil {
	//	// If no error, it means the event and user already exist
	//	return fmt.Errorf("user already joined this event")
	//} else if err != mongo.ErrNoDocuments {
	//	// If the error is not "no documents", return the error
	//	return fmt.Errorf("failed to check existing event: %w", err)
	//}
	//
	//// Proceed with further logic (e.g., insert the event) here
	//
	//result, err := r.eventsCollection.InsertOne(r.ctx, event)
	//if err != nil {
	//	return err
	//}
	////fmt.Println(result.InsertedID)
	//if result.InsertedID == nil {
	//	return errors.New("event not created")
	//}
	//return nil
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

	// Check Event

	count, err := r.eventsCollection.CountDocuments(r.ctx, bson.M{"eventId": eventId})
	if err != nil {
		return false, fmt.Errorf("failed to check event membership: %w", err)
	}
	if count == 0 {
		return false, errors.New("event not found")
	}
	// Check the join event of members
	filter := bson.M{
		"members": bson.M{
			"$elemMatch": bson.M{
				"eventId": eventId,
				"userId":  userId,
			},
		},
	}
	// check member is a join the event
	res := r.eventsCollection.FindOne(r.ctx, filter)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, res.Err()
	}
	var eventRes Event
	err = res.Decode(&eventRes)
	if err != nil {

		return false, err
	}
	members := eventRes.Members

	for _, member := range members {
		if member.UserId == userId {
			return true, nil
		}

	}
	return false, nil

	// Create a filter to check if the event exists for the given eventId and userId
	//eid := fmt.Sprintf("%s", eventId)
	//filter := bson.M{
	//	"eventId": eid,
	//	"userId":  userId,
	//}
	//
	//// Check if the event exists
	//count, err := r.eventsCollection.CountDocuments(r.ctx, filter)
	//if err != nil {
	//	return false, fmt.Errorf("failed to check event membership: %w", err)
	//}
	//
	//// If count is greater than 0, the user has joined the event
	//if count > 0 {
	//	return true, nil
	//}
	//return false, nil
}
func (r *eventRepositoryImpl) GetEventJoin(eventId string, userId string) (*MemberEventImpl, error) {

	// Check have the event
	countEvent, err := r.eventsCollection.CountDocuments(r.ctx, bson.M{"eventId": eventId})
	if err != nil {
		return nil, err
	}
	if countEvent == 0 {
		return nil, errors.New("event not found")
	}
	// Find event join data event members
	filter := bson.M{
		"members": bson.M{
			"$elemMatch": bson.M{
				"eventId": eventId,
				"userId":  userId,
			},
		}}

	// query data in events
	result := r.eventsCollection.FindOne(r.ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var eventRes Event
	err = result.Decode(&eventRes)
	if err != nil {
		return nil, err
	}
	members := eventRes.Members
	memberJoinInfo := MemberEventImpl{}
	for _, member := range members {
		memJoin := MemberEventImpl{
			EventId:        member.EventId,
			UserId:         member.UserId,
			JoinTime:       member.JoinTime,
			Name:           member.Name,
			LastName:       member.LastName,
			Organization:   member.Organization,
			Position:       member.Position,
			Course:         member.Course,
			LineId:         member.LineId,
			LineName:       member.LineName,
			ReferenceName:  member.ReferenceName,
			ReferencePhone: member.ReferencePhone,
			Clinic:         member.Clinic,
			Tel:            member.Tel,
		}
		memberJoinInfo = memJoin
		break
	}
	return &memberJoinInfo, nil

}
func (r *eventRepositoryImpl) CheckInEvent(userId string, eventCheckIn *EventCheckIn) (bool, error) {

	// find exits the event
	countEvent, err := r.eventsCollection.CountDocuments(r.ctx, bson.M{"eventId": eventCheckIn.EventId})
	if err != nil {
		return false, err
	}
	if countEvent == 0 {
		return false, errors.New("event not found")
	}
	log.Println("check in countEvent:", countEvent)
	// check memberCheckin and

	filter := bson.M{
		"eventCheckIn": bson.M{
			"$elemMatch": bson.M{
				"userId": userId,
				"checkIn": bson.M{
					"$exists": false,
				},
			},
		},
	}

	// Fine checkin for insert or update

	result := r.eventsCollection.FindOneAndUpdate(r.ctx, filter, bson.M{
		"$set": bson.M{
			"eventCheckIn": eventCheckIn,
		},
	}, options.FindOneAndUpdate().SetUpsert(false))

	if result.Err() != nil {

		if result.Err() == mongo.ErrNoDocuments {
			// insert memberCheckin to events
			updateResult, er := r.eventsCollection.UpdateOne(r.ctx, bson.M{"eventId": eventCheckIn.EventId}, bson.M{
				"$push": bson.M{
					"eventCheckIn": eventCheckIn,
				},
			})
			if er != nil {
				return false, er
			}
			if updateResult.ModifiedCount == 0 {
				return false, errors.New("event not found or no changes made")
			}
			return true, nil

		}
		return false, result.Err()
	}
	return true, nil

	//result := r.eventsCollection.FindOne(r.ctx, filter)
	//if result.Err() != nil {
	//	return false, result.Err()
	//}
	//err = result.Decode(&eventCheckIn)
	//if err != nil {
	//	return false, err
	//}

	//filer := bson.M{
	//	"eventId": eventCheckIn.EventId,
	//	"lineId":  userId,
	//}
	//
	//opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	//result := r.eventsCollection.FindOneAndUpdate(r.ctx, filer, bson.M{
	//	"$set": bson.M{
	//		"checkIn":      eventCheckIn.CheckIn,
	//		"checkOut":     eventCheckIn.CheckOut,
	//		"checkInTime":  eventCheckIn.CheckInTime,
	//		"checkOutTime": eventCheckIn.CheckOutTime,
	//		"checkInPlace": eventCheckIn.CheckInPlace,
	//	},
	//}, opts)
	//
	//if result.Err() != nil {
	//	return false, result.Err()
	//}
	//
	//return true, nil

}
func (r *eventRepositoryImpl) EventByUserId(userId string) ([]*EventResponse, error) {

	// get member join the event
	filter := bson.M{
		"members": bson.M{
			"$elemMatch": bson.M{
				"userId": userId,
			},
		},
	}
	// Opt
	//opts := options.FindOne().SetProjection(bson.M{"members": 1})
	//result := r.eventsCollection.FindOne(r.ctx, filter)
	results, err := r.eventsCollection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer results.Close(r.ctx)
	events := []*Event{}
	for results.Next(r.ctx) {
		var event Event
		err := results.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	eventRes := []*EventResponse{}
	for _, event := range events {
		isJoin := false
		for _, member := range event.Members {
			if member.UserId == userId {
				isJoin = true
				break
			}
		}
		item := EventResponse{
			EventId:          event.EventId,
			EventName:        event.Title,
			EventDescription: event.Description,
			EventStartDate:   time.Unix(event.StartDate, 0).Format("2006-01-02"),
			EventEndDate:     time.Unix(event.EndDate, 0).Format("2006-01-02"),
			EventPlace:       event.Place,
			EventStartTime:   time.Unix(event.EndTime, 0).Format("15:04"),
			EventBanner:      event.Banner,
			EventEndTime:     time.Unix(event.EndTime, 0).Format("15:04"),
			IsJoin:           isJoin,
		}

		eventRes = append(eventRes, &item)

	}
	return eventRes, nil
	//for _, event := range events {
	//	if event.Members != nil {
	//		for _, member := range event.Members {
	//			if member.UserId == userId {
	//
	//			}
	//		}
	//	}
	//}

	// members := eventRes.Members

	//Find event is not checkin filter by user Id
	//query := bson.M{
	//	"userId": userId,
	//	"$or": bson.A{
	//		bson.M{"checkIn": bson.M{"$exists": false}},
	//		bson.M{"checkIn": false},
	//	},
	//}
	//
	//events := []*MemberEventImpl{}
	//cursor, err := r.eventsCollection.Find(r.ctx, query)
	//if err != nil {
	//	return nil, err
	//}
	//defer cursor.Close(r.ctx)
	//for cursor.Next(r.ctx) {
	//	var event MemberEventImpl
	//	err := cursor.Decode(&event)
	//	if err != nil {
	//		return nil, err
	//	}
	//	events = append(events, &event)
	//}
	//count, _ := r.eventsCollection.CountDocuments(r.ctx, query)
	//log.Println(count)
	//return events, nil

}
func (r *eventRepositoryImpl) CreateEvent(event *Event) error {

	filter := bson.M{"eventId": event.EventId}

	exits, err := r.eventsCollection.CountDocuments(r.ctx, filter)
	if err != nil {
		return err
	}
	if exits > 0 {
		return errors.New("event already exists")
	}

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
	//backup members and eventCheckIn
	res := r.eventsCollection.FindOne(r.ctx, filter)
	if res.Err() != nil {
		return res.Err()
	}
	var eventRes Event
	err := res.Decode(&eventRes)
	if err != nil {
		return err
	}
	members := eventRes.Members
	event.Members = members

	// Create the update data using MongoDB's $set operator
	update := bson.M{
		"$set": event,
	}

	// Perform the update operation
	result, err := r.eventsCollection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}

	log.Println("result:", result)
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
