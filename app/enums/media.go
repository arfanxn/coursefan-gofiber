package enums

// models.Media.CollectionName enums
const (
	MediaCollectionNameUserAvatar       string = "user_avatar"
	MediaCollectionNameCourseImage      string = "course_image"
	MediaCollectionNameLectureVideo     string = "lecture_video"
	MediaCollectionNameLectureThumbnail string = "lecture_thumbnail"
	MediaCollectionNameDiscussion       string = "discussion"
	MediaCollectionNameMessage          string = "message"
)

// MediaCollectionNames returns slice of models.Media.CollectionName enums
func MediaCollectionNames() []string {
	return []string{
		MediaCollectionNameUserAvatar,
		MediaCollectionNameCourseImage,
		MediaCollectionNameLectureVideo,
		MediaCollectionNameLectureThumbnail,
		MediaCollectionNameDiscussion,
		MediaCollectionNameMessage,
	}
}
