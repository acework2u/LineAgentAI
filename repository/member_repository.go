package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
	
	return nil
}
func (r *memberRepositoryImpl) GetMember(id string) (*Member, error) {
	return nil, nil
}
func (r *memberRepositoryImpl) UpdateMember(member *Member) error {
	return nil
}
func (r *memberRepositoryImpl) DeleteMember(id string) error {
	return nil
}
func (r *memberRepositoryImpl) GetMembers() ([]*Member, error) {
	return nil, nil
}
