package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type LecturePartPolicy struct {
	repository *repositories.LecturePartRepository
}

func NewLecturePartPolicy(repository *repositories.LecturePartRepository) *LecturePartPolicy {
	return &LecturePartPolicy{
		repository: repository,
	}
}

func GetByCourse(c *fiber.Ctx) (err error) {
	return
}
