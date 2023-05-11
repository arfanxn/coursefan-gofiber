package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscussionRepository struct {
	db *gorm.DB
}

// NewDiscussionRepository instantiates a new DiscussionRepository
func NewDiscussionRepository(db *gorm.DB) *DiscussionRepository {
	return &DiscussionRepository{db: db}
}

// All returns all discussions in the database
func (repository *DiscussionRepository) All(c *fiber.Ctx, queries ...requests.Query) (
	discussions []models.Discussion, err error) {
	tx := repository.db
	if query := sliceh.FirstOrNil(queries); query != nil {
		tx = gormh.BuildFromRequestQuery(repository.db, models.Discussion{}, *query)
	}
	err = tx.Find(&discussions).Error
	return
}

// AllByLecture returns all discussion by lecture
func (repository *DiscussionRepository) AllByLecture(c *fiber.Ctx, query requests.Query) (
	discussions []models.Discussion, err error) {
	lectureIdFilter := query.GetFilter(
		models.Discussion{}.TableName()+".discussable_id",
		enums.QueryFilterOperatorEquals)
	err = gormh.BuildFromRequestQuery(repository.db, models.Discussion{}, query).
		Where(models.Discussion{}.TableName()+".discussable_type = ?", reflecth.GetTypeName(models.Lecture{})).
		Where(models.Discussion{}.TableName()+".discussable_id = ?", lectureIdFilter.Values[0]).
		Distinct().Find(&discussions).Error
	return
}

// Find finds model by id
func (repository *DiscussionRepository) Find(c *fiber.Ctx, id string) (review models.Discussion, err error) {
	err = repository.db.Where("id = ?", id).First(&review).Error
	return
}

// FindById finds model by id
func (repository *DiscussionRepository) FindById(c *fiber.Ctx, id string) (review models.Discussion, err error) {
	err = repository.db.Where("id = ?", id).First(&review).Error
	return
}

// Insert inserts models into the database
func (repository *DiscussionRepository) Insert(c *fiber.Ctx, discussions ...*models.Discussion) (int64, error) {
	for _, discussion := range discussions {
		if discussion.Id == uuid.Nil {
			discussion.Id = uuid.New()
		}
	}
	result := repository.db.Create(discussions)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *DiscussionRepository) UpdateById(c *fiber.Ctx, discussion *models.Discussion) (int64, error) {
	result := repository.db.Model(discussion).Updates(discussion)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *DiscussionRepository) DeleteByIds(c *fiber.Ctx, discussions ...*models.Discussion) (int64, error) {
	result := repository.db.Delete(discussions)
	return result.RowsAffected, result.Error
}
