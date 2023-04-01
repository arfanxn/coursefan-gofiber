package enums

// models.Role.Name enums
const (
	// Course module
	RoleNameCourseLecturer    string = "course_lecturer"
	RoleNameCourseParticipant string = "course_participant"
)

// RoleNames returns slice of models.Role.Name enums
func RoleNames() []string {
	return []string{
		RoleNameCourseLecturer,
		RoleNameCourseParticipant,
	}
}
