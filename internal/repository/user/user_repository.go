package user

import (
	"context"
	"errors"

	"github.com/dakasakti/deploy-apps-hexagonal/internal/middlewares"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type UserRepository interface {
	Gets(ctx context.Context) ([]models.User, error)
	Get(ctx context.Context, user_id uint) (*models.User, error)
	Create(ctx context.Context, data models.User) (interface{}, error)
	Update(ctx context.Context, data models.User) error
	Delete(ctx context.Context, user_id uint) error
	Login(data models.User) (*string, error)
}

type userRepository struct {
	db *gorm.DB
	mc *mongo.Collection
}

func NewUserRepository(db *gorm.DB, mc *mongo.Collection) *userRepository {
	return &userRepository{
		db: db,
		mc: mc,
	}
}

func (ur *userRepository) Gets(ctx context.Context) ([]models.User, error) {
	var users []models.User

	results, err := ur.mc.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		err := results.Decode(&singleUser)
		if err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}

	return users, nil
}

func (ur *userRepository) Get(ctx context.Context, user_id uint) (*models.User, error) {
	var user models.User

	err := ur.mc.FindOne(ctx, bson.M{"id": user_id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Create(ctx context.Context, data models.User) (interface{}, error) {
	err := ur.db.Create(&data).Error
	if err != nil {
		return nil, err
	}

	result, err := ur.mc.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (ur *userRepository) Update(ctx context.Context, data models.User) error {
	err := ur.db.Updates(&data).Error
	if err != nil {
		return err
	}

	update := bson.M{"fullname": data.Fullname, "username": data.Username, "password": data.Password}
	_, err = ur.mc.UpdateOne(ctx, bson.M{"id": data.ID}, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Delete(ctx context.Context, user_id uint) error {
	err := ur.db.Delete(&models.User{}, user_id).Error
	if err != nil {
		return err
	}

	_, err = ur.mc.DeleteOne(ctx, bson.M{"id": user_id})
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Login(data models.User) (*string, error) {
	tx := ur.db.Find(&data, "username = ? AND password = ?", data.Username, data.Password)
	if tx.RowsAffected != 1 {
		return nil, errors.New("username atau password salah")
	}

	token, err := middlewares.CreateToken(data.ID, data.Username)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
