package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CoursePolicy struct {
	curRepository *repositories.CourseUserRoleRepository
}

func NewCoursePolicy(curRepository *repositories.CourseUserRoleRepository) *CoursePolicy {
	return &CoursePolicy{curRepository: curRepository}
}

func (policy *CoursePolicy) Update(c *fiber.Ctx, input requests.CourseUpdate) (err error) {
	curMdl, err := policy.curRepository.FindByModel(c, models.CourseUserRole{
		CourseId: uuid.MustParse(input.Id),
		UserId:   ctxh.MustGetUser(c).Id,
		Relation: enums.CourseUserRoleRelationLecturer,
	})
	if errorh.IsGormErrRecordNotFound(err) || curMdl.Id == uuid.Nil {
		err = fiber.ErrForbidden
		return
	} else if err != nil {
		return
	}
	return nil
}

func (policy *CoursePolicy) Delete(c *fiber.Ctx, input requests.CourseDelete) (err error) {
	curMdl, err := policy.curRepository.FindByModel(c, models.CourseUserRole{
		CourseId: uuid.MustParse(input.Id),
		UserId:   ctxh.MustGetUser(c).Id,
		Relation: enums.CourseUserRoleRelationLecturer,
	})
	if errorh.IsGormErrRecordNotFound(err) || curMdl.Id == uuid.Nil {
		err = fiber.ErrForbidden
		return
	} else if err != nil {
		return
	}
	return nil
}
