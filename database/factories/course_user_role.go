package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

func FakeCourseUserRole() models.CourseUserRole {
	return models.CourseUserRole{
		// Id:, will be filled in later
		// CourseId:, // will be filled in later
		// Course:, // will be filled in later
		// UserId:, // will be filled in later
		// User:, // will be filled in later
		// RoleId:, // will be filled in later
		// Role:, // will be filled in later
		Relation:  sliceh.Random(enums.CourseUserRoleRelations()...), // will be filled in later
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
