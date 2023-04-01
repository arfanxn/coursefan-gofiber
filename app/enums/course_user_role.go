package enums

// models.CourseUserRole.Relation enums
const (
	CourseUserRoleRelationCart        string = "course_in_cart"
	CourseUserRoleRelationLecturer    string = "course_lecturer"
	CourseUserRoleRelationParticipant string = "course_participant"
	CourseUserRoleRelationWishlist    string = "course_in_wishlist"
)

// CourseUserRoleRelations returns slice of models.CourseUserRole.Relation enums
func CourseUserRoleRelations() []string {
	return []string{
		CourseUserRoleRelationCart,
		CourseUserRoleRelationLecturer,
		CourseUserRoleRelationParticipant,
		CourseUserRoleRelationWishlist,
	}
}
