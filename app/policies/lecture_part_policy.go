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

// / AllByCourse policy ensures that the user has the right permissions for access a lecture parts.
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
		// allow course participants and lecturer to access the lecture parts
		enums.CourseUserRoleRelationParticipant,
		enums.CourseUserRoleRelationLecturer,
	}, curMdl.Relation) == false {
		err = fiber.ErrForbidden
		return
	}
	return nil
}

// Find policy ensures that the user has the right permissions for access a lecture part.
func (policy *LecturePartPolicy) Find(c *fiber.Ctx, input requests.Query) (err error) {
	lecturePartId, courseId := c.Params("lecture_part_id"), c.Params("course_id")
	curMdl, err := policy.curRepository.FindByModel(c, models.CourseUserRole{
		CourseId: uuid.MustParse(courseId),
		UserId:   ctxh.MustGetUser(c).Id,
	})
	if errorh.IsGormErrRecordNotFound(err) || curMdl.Id == uuid.Nil {
		err = fiber.ErrForbidden
		return
	} else if err != nil {
		return
	} else if sliceh.Contains([]string{
		// allow course participants and lecturer to access the lecture part
		enums.CourseUserRoleRelationParticipant,
		enums.CourseUserRoleRelationLecturer,
	}, curMdl.Relation) == false {
		err = fiber.ErrForbidden
		return
	}
	lecturePartMdl, err := policy.repository.FindByModel(c, models.LecturePart{Id: uuid.MustParse(lecturePartId)})
	if errorh.IsGormErrRecordNotFound(err) || lecturePartMdl.Id == uuid.Nil {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	} else if curMdl.CourseId != lecturePartMdl.CourseId {
		err = fiber.ErrForbidden
		return
	}
	return nil
}

// Create policy ensures that the user has the right permissions for create a lecture part.
func (policy *LecturePartPolicy) Create(c *fiber.Ctx, input requests.LecturePartCreate) (err error) {
	curMdl, err := policy.curRepository.FindByModel(c, models.CourseUserRole{
		CourseId: uuid.MustParse(input.CourseId),
		UserId:   ctxh.MustGetUser(c).Id,
	})
	if errorh.IsGormErrRecordNotFound(err) || curMdl.Id == uuid.Nil {
		err = fiber.ErrForbidden
		return
	} else if err != nil {
		return
	} else if sliceh.Contains([]string{
		// allow course lecturer to create lecture part
		enums.CourseUserRoleRelationLecturer,
	}, curMdl.Relation) == false {
		err = fiber.ErrForbidden
		return
	}
	return nil
}

// Update policy ensures that the user has the right permissions for update a lecture part.
func (policy *LecturePartPolicy) Update(c *fiber.Ctx, input requests.LecturePartUpdate) (err error) {
	curMdl, err := policy.curRepository.FindByModel(c, models.CourseUserRole{
		CourseId: uuid.MustParse(input.CourseId),
		UserId:   ctxh.MustGetUser(c).Id,
	})
	if errorh.IsGormErrRecordNotFound(err) || curMdl.Id == uuid.Nil {
		err = fiber.ErrForbidden
		return
	} else if err != nil {
		return
	} else if sliceh.Contains([]string{
		// allow course lecturer to update lecture part
		enums.CourseUserRoleRelationLecturer,
	}, curMdl.Relation) == false {
		err = fiber.ErrForbidden
		return
	}
	lecturePartMdl, err := policy.repository.FindByModel(c, models.LecturePart{Id: uuid.MustParse(input.Id)})
	if errorh.IsGormErrRecordNotFound(err) || lecturePartMdl.Id == uuid.Nil {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	} else if curMdl.CourseId != lecturePartMdl.CourseId {
		err = fiber.ErrForbidden
		return
	}
	return nil
}

// Delete policy ensures that the user has the right permissions for delete a lecture part.
func (policy *LecturePartPolicy) Delete(c *fiber.Ctx, input requests.LecturePartDelete) (err error) {
	curMdl, err := policy.curRepository.FindByModel(c, models.CourseUserRole{
		CourseId: uuid.MustParse(input.CourseId),
		UserId:   ctxh.MustGetUser(c).Id,
	})
	if errorh.IsGormErrRecordNotFound(err) || curMdl.Id == uuid.Nil {
		err = fiber.ErrForbidden
		return
	} else if err != nil {
		return
	} else if sliceh.Contains([]string{
		// allow course lecturer to delete lecture part
		enums.CourseUserRoleRelationLecturer,
	}, curMdl.Relation) == false {
		err = fiber.ErrForbidden
		return
	}
	lecturePartMdl, err := policy.repository.FindByModel(c, models.LecturePart{Id: uuid.MustParse(input.Id)})
	if errorh.IsGormErrRecordNotFound(err) || lecturePartMdl.Id == uuid.Nil {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	} else if curMdl.CourseId != lecturePartMdl.CourseId {
		err = fiber.ErrForbidden
		return
	}
	return nil
}
