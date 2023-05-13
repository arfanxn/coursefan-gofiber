package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
)

type CourseOrderPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewCourseOrderPolicy(
	permissionRepository *repositories.PermissionRepository,
) *CourseOrderPolicy {
	return &CourseOrderPolicy{
		permissionRepository: permissionRepository,
	}
}
