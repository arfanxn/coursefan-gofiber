package enums

// models.Permission.Name enums
const (
	// Notification module
	PermissionNameNotificationView             string = "notification.view"
	PermissionNameNotificationCreate           string = "notification.create"
	PermissionNameNotificationEdit             string = "notification.edit"
	PermissionNameNotificationMarkReadOrUnread string = "notification.mark_read_or_unread"
	PermissionNameNotificationDelete           string = "notification.delete"

	// Course module
	PermissionNameCourseView   string = "course.view"
	PermissionNameCourseCreate string = "course.create"
	PermissionNameCourseEdit   string = "course.edit"
	PermissionNameCourseDelete string = "course.delete"

	// Course Order module
	PermissionNameCourseOrderView   string = "course_order.view"
	PermissionNameCourseOrderCreate string = "course_order.create"
	PermissionNameCourseOrderEdit   string = "course_order.edit"
	PermissionNameCourseOrderDelete string = "course_order.delete"

	// Course Member/User module
	PermissionNameCourseUserView   string = "course_user.view"
	PermissionNameCourseUserCreate string = "course_user.create"
	PermissionNameCourseUserEdit   string = "course_user.edit"
	PermissionNameCourseUserDelete string = "course_user.delete"

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
	PermissionNameDiscussionUpvote string = "lecture.upvote"
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
		PermissionNameNotificationView,
		PermissionNameNotificationCreate,
		PermissionNameNotificationEdit,
		PermissionNameNotificationMarkReadOrUnread,
		PermissionNameNotificationDelete,

		PermissionNameCourseView,
		PermissionNameCourseCreate,
		PermissionNameCourseEdit,
		PermissionNameCourseDelete,

		PermissionNameCourseOrderView,
		PermissionNameCourseOrderCreate,
		PermissionNameCourseOrderEdit,
		PermissionNameCourseOrderDelete,

		PermissionNameCourseUserView,
		PermissionNameCourseUserCreate,
		PermissionNameCourseUserEdit,
		PermissionNameCourseUserDelete,

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
		PermissionNameDiscussionUpvote,
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
