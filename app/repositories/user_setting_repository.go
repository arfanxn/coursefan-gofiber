package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type UserSettingRepository struct {
	db *gorm.DB
}

// NewUserSettingRepository instantiates a new UserSettingRepository
func NewUserSettingRepository(db *gorm.DB) *UserSettingRepository {
	return &UserSettingRepository{db: db}
}

// All returns all user settings in the database
func (repository *UserSettingRepository) All(c *fiber.Ctx, queries ...requests.Query) (
	uss []models.UserSetting, err error) {
	tx := repository.db
	if query := sliceh.FirstOrNil(queries); query != nil {
		tx = gormh.BuildFromRequestQuery(repository.db, models.UserSetting{}, *query)
	}
	err = tx.Find(&uss).Error
	return
}

// Find finds model by id
func (repository *UserSettingRepository) FindById(c *fiber.Ctx, id string) (user models.UserSetting, err error) {
	err = repository.db.Where("id = ?", id).First(&user).Error
	return
}

// Find finds model by model
func (repository *UserSettingRepository) FindByModel(c *fiber.Ctx, model models.UserSetting) (
	us models.UserSetting, err error) {
	err = repository.db.First(&us, model).Error
	return
}

// Insert inserts models into the database
func (repository *UserSettingRepository) Insert(c *fiber.Ctx, userSettings ...*models.UserSetting) (int64, error) {
	for _, userSetting := range userSettings {
		if userSetting.Id == uuid.Nil {
			userSetting.Id = uuid.New()
		}
		userSetting.CreatedAt = time.Now()
	}
	result := repository.db.Create(userSettings)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *UserSettingRepository) UpdateById(c *fiber.Ctx, userSetting *models.UserSetting) (int64, error) {
	// refresh model updated at
	userSetting.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(userSetting).Updates(userSetting)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *UserSettingRepository) DeleteByIds(c *fiber.Ctx, userSettings ...*models.UserSetting) (int64, error) {
	result := repository.db.Delete(userSettings)
	return result.RowsAffected, result.Error
}
