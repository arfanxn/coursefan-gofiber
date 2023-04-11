package enums

// models.Course query scope enums
const (
	CourseQueryScopeLectured     string = "lecturerd"
	CourseQueryScopeParticipated string = "participated"
	CourseQueryScopeWishlist     string = "wishlist"
	CourseQueryScopeCart       string = "in_cart"
)

// CourseQueryScopes returns slice of Course query scope enums
func CourseQueryScopes() []string {
	return []string{
		CourseQueryScopeLectured,
		CourseQueryScopeParticipated,
		CourseQueryScopeWishlist,
		CourseQueryScopeCart,
	}
}
