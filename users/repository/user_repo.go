package repository

import (
	"context"
	"github.com/enrinal/demo-order-go/constant"
	"github.com/enrinal/demo-order-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Store(ctx context.Context, user models.User) error
}

func NewUserRepo(mgo *mongo.Client) *UserRepo {
	return &UserRepo{mgo}
}

type UserRepo struct {
	mgo *mongo.Client
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	collection := r.mgo.Database(constant.MONGO_DATABASE_NAME).Collection(constant.COLLECTION_USERS)
	err := collection.FindOne(ctx, map[string]interface{}{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) Store(ctx context.Context, user models.User) error {
	collection := r.mgo.Database(constant.MONGO_DATABASE_NAME).Collection(constant.COLLECTION_USERS)
	user.ID = primitive.NewObjectID()
	user.CreateAt = time.Now().String()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
