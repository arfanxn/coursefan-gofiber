package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LecturePartPolicy struct {
	repository    *repositories.LecturePartRepository
	curRepository *repositories.CourseUserRoleRepository
}

func NewLecturePartPolicy(
	repository *repositories.LecturePartRepository,
	curRepository *repositories.CourseUserRoleRepository,
) *LecturePartPolicy {
	return &LecturePartPolicy{
		repository:    repository,
		curRepository: curRepository,
	}
}

func (policy *LecturePartPolicy) AllByCourse(c *fiber.Ctx, input requests.Query) (err error) {
	curMdl, err := policy.curRepository.FindByModel(c, models.CourseUserRole{
		CourseId: uuid.MustParse(c.Params("course_id")),
		UserId:   ctxh.MustGetUser(c).Id,
	})
	if errorh.IsGormErrRecordNotFound(err) || curMdl.Id == uuid.Nil {
		err = fiber.ErrForbidden
		return
	} else if err != nil {
		return
	} else if sliceh.Contains([]string{
		enums.CourseUserRoleRelationParticipant,
		enums.CourseUserRoleRelationLecturer,
	}, curMdl.Relation) == false {
		err = fiber.ErrForbidden
		return
	}
	return nil
}
