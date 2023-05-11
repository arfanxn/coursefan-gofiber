package enums

// models.Permission.Name enums
const (
	// Course module
	PermissionNameCourseView   string = "course.view"
	PermissionNameCourseCreate string = "course.create"
	PermissionNameCourseEdit   string = "course.edit"
	PermissionNameCourseDelete string = "course.delete"

	// Course Review module
	PermissionNameCourseReviewView   string = "course_review.view"
	PermissionNameCourseReviewCreate string = "course_review.create"
	PermissionNameCourseReviewEdit   string = "course_review.edit"
	PermissionNameCourseReviewDelete string = "course_review.delete"

	// Review module
	PermissionNameReviewView   string = "review.view"
	PermissionNameReviewCreate string = "review.create"
	PermissionNameReviewEdit   string = "review.edit"
	PermissionNameReviewDelete string = "review.delete"

	// Lecture module
	PermissionNameLectureView   string = "lecture.view"
	PermissionNameLectureCreate string = "lecture.create"
	PermissionNameLectureEdit   string = "lecture.edit"
	PermissionNameLectureDelete string = "lecture.delete"

	// Lecture Part module
	PermissionNameLecturePartView   string = "lecture_part.view"
	PermissionNameLecturePartCreate string = "lecture_part.create"
	PermissionNameLecturePartEdit   string = "lecture_part.edit"
	PermissionNameLecturePartDelete string = "lecture_part.delete"

	// Discussion module
	PermissionNameDiscussionView   string = "lecture.view"
	PermissionNameDiscussionCreate string = "lecture.create"
	PermissionNameDiscussionEdit   string = "lecture.edit"
	PermissionNameDiscussionDelete string = "lecture.delete"

	// Lecture Discussion module
	PermissionNameLectureDiscussionView   string = "lecture_discussion.view"
	PermissionNameLectureDiscussionCreate string = "lecture_discussion.create"
	PermissionNameLectureDiscussionEdit   string = "lecture_discussion.edit"
	PermissionNameLectureDiscussionDelete string = "lecture_discussion.delete"

	// Lecture Progress module
	PermissionNameLectureProgressView   string = "lecture_progress.view"
	PermissionNameLectureProgressCreate string = "lecture_progress.create"
	PermissionNameLectureProgressEdit   string = "lecture_progress.edit"
	PermissionNameLectureProgressDelete string = "lecture_progress.delete"
)

// PermissionNames returns slice of models.Permission.Name enums
func PermissionNames() []string {
	return []string{
		PermissionNameCourseView,
		PermissionNameCourseCreate,
		PermissionNameCourseEdit,
		PermissionNameCourseDelete,

		PermissionNameCourseReviewView,
		PermissionNameCourseReviewCreate,
		PermissionNameCourseReviewEdit,
		PermissionNameCourseReviewDelete,

		PermissionNameReviewView,
		PermissionNameReviewCreate,
		PermissionNameReviewEdit,
		PermissionNameReviewDelete,

		PermissionNameLectureView,
		PermissionNameLectureCreate,
		PermissionNameLectureEdit,
		PermissionNameLectureDelete,

		PermissionNameLecturePartView,
		PermissionNameLecturePartCreate,
		PermissionNameLecturePartEdit,
		PermissionNameLecturePartDelete,

		PermissionNameDiscussionView,
		PermissionNameDiscussionCreate,
		PermissionNameDiscussionEdit,
		PermissionNameDiscussionDelete,

		PermissionNameLectureDiscussionView,
		PermissionNameLectureDiscussionCreate,
		PermissionNameLectureDiscussionEdit,
		PermissionNameLectureDiscussionDelete,

		PermissionNameLectureProgressView,
		PermissionNameLectureProgressCreate,
		PermissionNameLectureProgressEdit,
		PermissionNameLectureProgressDelete,
	}

}
