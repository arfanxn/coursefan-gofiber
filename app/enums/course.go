package enums

// models.Course query scope enums
const (
	CourseQueryScopeLectured     string = "lectured"
	CourseQueryScopeParticipated string = "participated"
	CourseQueryScopeWishlist     string = "wishlist"
	CourseQueryScopeCart         string = "cart"
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
