package enums

// models.Role.Name enums
const (
	RoleNameCourseLecturer    string = "course_lecturer"
	RoleNameCourseParticipant string = "course_participant"
	RoleNameCourseWishlister  string = "course_wishlister"
	RoleNameCourseCarter      string = "course_carter"
)

// RoleNames returns slice of models.Role.Name enums
func RoleNames() []string {
	return []string{
		RoleNameCourseLecturer,
		RoleNameCourseParticipant,
		RoleNameCourseWishlister,
		RoleNameCourseCarter,
	}
}
