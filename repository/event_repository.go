package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
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

	// Define the filter to find the specific member in the event
	memberFilter := bson.M{
		"eventId":        event.EventId,
		"members.userId": event.UserId,
	}

	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"members.$[member].eventId":        event.EventId,
			"members.$[member].joinTime":       event.JoinTime,
			"members.$[member].name":           event.Name,
			"members.$[member].lastName":       event.LastName,
			"members.$[member].organization":   event.Organization,
			"members.$[member].position":       event.Position,
			"members.$[member].course":         event.Course,
			"members.$[member].lineId":         event.LineId,
			"members.$[member].lineName":       event.LineName,
			"members.$[member].tel":            event.Tel,
			"members.$[member].referenceName":  event.ReferenceName,
			"members.$[member].referencePhone": event.ReferencePhone,
			"members.$[member].clinic":         event.Clinic,
		},
	}

	// Define array filters for the update
	opts := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"member.userId": event.UserId},
		},
	})

	// Try to update the existing member
	updateResult, err := r.eventsCollection.UpdateOne(r.ctx, memberFilter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to update member: %w", err)
	}

	// If no document was modified, insert the member
	if updateResult.ModifiedCount == 0 {
		// Use $push to add the new member
		insertUpdate := bson.M{
			"$push": bson.M{
				"members": event,
			},
		}
		_, err := r.eventsCollection.UpdateOne(r.ctx, bson.M{"eventId": event.EventId}, insertUpdate)
		if err != nil {
			return fmt.Errorf("failed to insert new member: %w", err)
		}
	}

	return nil

}
func (r *eventRepositoryImpl) EventLeave(event *MemberEventImpl) error {
	panic("implement me")
}
func (r *eventRepositoryImpl) GetEvent(eventId string) (*MemberEventImpl, error) {
	panic("implement me")
}
func (r *eventRepositoryImpl) GetEvents(filter EventFilter) ([]*MemberEventImpl, error) {
	panic("implement me")
}
func (r *eventRepositoryImpl) CheckJoinEvent(eventId string, userId string) (bool, error) {

	// Check Event
	log.Println("eventId:", eventId)
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
	log.Println("res:", res)
	if res.Err() != nil {

		log.Println("res.Err():", res.Err())

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

	// last member is check in
	checkin := []*EventCheckIn{}
	for _, eventCheckIn := range eventRes.EventCheckIn {
		if eventCheckIn.UserId == userId {
			checkin = append(checkin, eventCheckIn)
		}
	}

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
			EventCheckIn:   checkin,
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
func (r *eventRepositoryImpl) EventByUserId(userId string) ([]*Event, error) {

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

	eventRes := []*Event{}
	myActiveEvent := []MemberEventImpl{}
	for _, event := range events {
		isJoin := false
		for _, member := range event.Members {
			if member.UserId == userId {
				isJoin = true
				myInfo := MemberEventImpl{
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
					Tel:            member.Tel,
					ReferenceName:  member.ReferenceName,
					ReferencePhone: member.ReferencePhone,
					Clinic:         member.Clinic,
				}

				myActiveEvent = append(myActiveEvent, myInfo)
				break
			}
		}
		item := Event{
			EventId:     event.EventId,
			Title:       event.Title,
			Description: event.Description,
			StartDate:   event.StartDate,
			EndDate:     event.EndDate,
			Place:       event.Place,
			StartTime:   event.StartTime,
			Banner:      event.Banner,
			EndTime:     event.EndTime,
			Location:    event.Location,
			Status:      isJoin,
			CreatedDate: event.CreatedDate,
			UpdatedDate: event.UpdatedDate,
			LineId:      event.LineId,
			LineName:    event.LineName,
			EventType:   event.EventType,
			Members:     myActiveEvent,
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
func (r *eventRepositoryImpl) EventsList(filter EventFilter) ([]*Event, error) {

	log.Println("filter in repo:", filter)
	events := []*Event{}
	// startDate stage
	// check filter start date
	pipeline := []bson.M{}
	startDateStage := bson.M{}
	endDateStage := bson.M{}

	if filter.Start > 0 {
		startDateStage = bson.M{
			"$match": bson.M{
				"startDate": bson.M{
					"$gte": filter.Start,
				},
			},
		}
		pipeline = append(pipeline, startDateStage)
	}

	if filter.End > 0 {
		endDateStage = bson.M{
			"$match": bson.M{
				"endDate": bson.M{
					"$lte": filter.End,
				},
			},
		}
		pipeline = append(pipeline, endDateStage)
	}

	// and match status is true or false
	//log.Println("filter Stage:", filter.Stages)

	statusStage := bson.M{
		"$match": bson.M{
			"status": filter.Status,
		},
	}
	if strings.Compare(filter.Stages, "all") == 0 {
		//log.Println("filter Stage is:", filter.Stages)
		// status is true or false
		statusStage = bson.M{
			"$match": bson.M{
				"status": bson.M{
					"$exists": true,
					"$type":   "bool",
				},
			},
		}

	}
	//
	//if strings.Compare(filter.Stages, "all") == 0 {
	//	log.Println("filter Stage is:", filter.Stages)
	//
	//}

	//limit stage
	if filter.Limit > 0 {
		limitStage := bson.M{
			"$limit": filter.Limit,
		}
		pipeline = append(pipeline, limitStage)
	}
	// sort stage
	if filter.Sort == "asc" {
		sortStage := bson.M{
			"$sort": bson.M{
				"startDate": 1,
			},
		}
		pipeline = append(pipeline, sortStage)
	} else if filter.Sort == "desc" {
		sortStage := bson.M{
			"$sort": bson.M{
				"startDate": -1,
			},
		}
		pipeline = append(pipeline, sortStage)
	}

	pipeline = append(pipeline, statusStage)

	log.Println("pipeline:", pipeline, statusStage)
	//pipeline = append(pipeline, startDateStage, endDateStage)

	curr, err := r.eventsCollection.Aggregate(r.ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer curr.Close(r.ctx)
	for curr.Next(r.ctx) {
		var event Event
		err := curr.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil

	// order evnet by id desc
	//opts := options.Find().SetSort(bson.D{{"_id", -1}})
	//cursor, err := r.eventsCollection.Find(r.ctx, filter, opts)
	//
	//if err != nil {
	//	return nil, err
	//}
	//defer cursor.Close(r.ctx)
	//for cursor.Next(r.ctx) {
	//	var event Event
	//	err := cursor.Decode(&event)
	//	if err != nil {
	//		return nil, err
	//	}
	//	events = append(events, &event)
	//}
	//return events, nil
}
func (r *eventRepositoryImpl) EventsByClinic(eventId string) ([]*ClinicGroup, error) {

	// Define the aggregate pipe line
	pipeline := []bson.M{
		{"$match": bson.M{"eventId": eventId}}, {
			"$unwind": "$members",
		}, {
			"$group": bson.M{
				"_id":     "$members.clinic",
				"members": bson.M{"$push": "$members"},
				"count":   bson.M{"$sum": 1},
			},
		},
		{"$sort": bson.M{"_id": 1}},
	}

	cursor, err := r.eventsCollection.Aggregate(r.ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)
	clinics := []*ClinicGroup{}
	for cursor.Next(r.ctx) {
		var clinic ClinicGroup
		err := cursor.Decode(&clinic)
		if err != nil {
			return nil, err
		}
		clinics = append(clinics, &clinic)
	}

	return clinics, nil
}
func (r *eventRepositoryImpl) EventReport(filter *ReportFilter) ([]*Event, error) {
	// Define the pipeline with a match stage to filter by startDate and endDate
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"startDate": bson.M{
					"$gte": filter.StartDate,
					"$lte": filter.EndDate,
				},
			},
		},
		{
			"$sort": bson.M{"startDate": 1}, // Sort by startDate in ascending order
		},
	}

	cursor, err := r.eventsCollection.Aggregate(r.ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	// Decode the results into a slice of EventResponse
	events := []*Event{}
	for cursor.Next(r.ctx) {
		var event Event
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
func (r *eventRepositoryImpl) CountEvent(filter EventFilter) (int, error) {
	// count all event with filter
	query := bson.M{}
	if filter.Start > 0 {
		query["startDate"] = bson.M{
			"$gte": filter.Start,
		}
	}
	if filter.End > 0 {
		query["endDate"] = bson.M{
			"$lte": filter.End,
		}
	}
	if filter.Status {
		query["status"] = filter.Status
	}
	if filter.Keyword != "" {
		query["$text"] = bson.M{
			"$search": filter.Keyword,
		}
	}
	if filter.Sort == "asc" {
		query["startDate"] = 1
	} else if filter.Sort == "desc" {
		query["startDate"] = -1
	}

	count, err := r.eventsCollection.CountDocuments(r.ctx, query)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
func (r *eventRepositoryImpl) CountMemberJoinEvents(filter EventFilter) (int, error) {
	query := bson.M{}
	if filter.Start > 0 {
		query["startDate"] = bson.M{
			"$gte": filter.Start,
		}
	}
	if filter.End > 0 {
		query["endDate"] = bson.M{
			"$lte": filter.End,
		}
	}
	if filter.Status {
		query["status"] = filter.Status
	}
	if filter.Keyword != "" {
		query["$text"] = bson.M{
			"$search": filter.Keyword,
		}
	}
	if filter.Sort == "asc" {
		query["startDate"] = 1
	}
	if filter.Sort == "desc" {
		query["startDate"] = -1
	}
	if filter.Stages == "all" {
		query["status"] = bson.M{
			"$exists": true,
			"$type":   "bool",
		}
	}
	if filter.Stages == "active" {
		query["status"] = true
	}
	if filter.Stages == "inactive" {
		query["status"] = false
	}
	limit := 10
	if filter.Limit > 0 {
		limit = filter.Limit
	}
	_ = limit

	_ = query
	pipeline := []bson.M{}
	pipeline = append(pipeline, bson.M{"$match": query})
	pipeline = append(pipeline, bson.M{"$unwind": "$members"})
	// pipeline group ny _id is null
	//pipeline = append(pipeline, bson.M{"$group": bson.M{"_id": "null", "count": bson.M{"$sum": 1}}})
	pipeline = append(pipeline, bson.M{"$group": bson.M{"_id": "members.userId", "allMembers": bson.M{"$addToSet": "$members"}}})
	pipeline = append(pipeline, bson.M{"$unwind": "$allMembers"})
	//pipeline = append(pipeline, bson.M{"$sort": bson.M{"count": -1}})
	//pipeline = append(pipeline, bson.M{"$limit": limit})
	cursor, err := r.eventsCollection.Aggregate(r.ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(r.ctx)
	count := 0
	for cursor.Next(r.ctx) {
		var event Event
		err := cursor.Decode(&event)
		if err != nil {
			return 0, err
		}
		count++
	}
	if err = cursor.Err(); err != nil {
		return 0, err
	}
	return count, nil

}
