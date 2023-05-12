package repositories

import (
	"fmt"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository instantiates a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// All returns all users in the database
func (repository *UserRepository) All(c *fiber.Ctx, queries ...requests.Query) (users []models.User, err error) {
	tx := repository.db
	if query := sliceh.FirstOrNil(queries); query != nil {
		tx = gormh.BuildFromRequestQuery(repository.db, models.Review{}, *query)
	}
	err = tx.Find(&users).Error
	return
}

// AllByCourse returns all users by course
func (repository *UserRepository) AllByCourse(c *fiber.Ctx, query requests.Query) (
	users []models.User, err error) {
	courseIdFilter := query.GetFilter(models.Course{}.TableName()+".id", enums.QueryFilterOperatorEquals)
	err = gormh.BuildFromRequestQuery(repository.db, models.User{}, query).
		Joins(
			fmt.Sprintf("JOIN %s ON %s.%s = %s.%s",
				models.CourseUserRole{}.TableName(),
				models.CourseUserRole{}.TableName(),
				"user_id",
				models.User{}.TableName(),
				"id",
			)).
		Joins(
			fmt.Sprintf("JOIN %s ON %s.%s = %s.%s",
				models.Course{}.TableName(),
				models.Course{}.TableName(),
				"id",
				models.CourseUserRole{}.TableName(),
				"course_id",
			)).
		Where(models.Course{}.TableName()+".id = ?", courseIdFilter.Values[0]).
		Distinct().Find(&users).Error

	return
}

// FindByEmail finds a user by email
func (repository *UserRepository) FindByEmail(c *fiber.Ctx, email string) (user models.User, err error) {
	err = repository.db.Where("email = ?", email).First(&user).Error
	return
}

// Insert inserts users into the database
func (repository *UserRepository) Insert(c *fiber.Ctx, users ...*models.User) (int64, error) {
	for _, user := range users {
		if user.Id == uuid.Nil {
			user.Id = uuid.New()
		}
		user.CreatedAt = time.Now()
	}
	result := repository.db.Create(users)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *UserRepository) UpdateById(c *fiber.Ctx, user *models.User) (int64, error) {
	// refresh model updated at
	user.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(user).Updates(user)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *UserRepository) DeleteByIds(c *fiber.Ctx, users ...*models.User) (int64, error) {
	result := repository.db.Delete(users)
	return result.RowsAffected, result.Error
}
