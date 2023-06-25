package admin_dto

type CourseIDParam struct {
	ID int `json:"id"`
}

type CourseLessonIDParam struct {
	CourseID int `json:"course_id"`
	LessonID int `json:"lesson_id"`
}

type CourseAllRes struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Desc   string   `json:"desc"`
	Author []string `json:"author"`
}

type CourseRes struct {
	Title       string   `json:"title"`
	Desc        string   `json:"desc"`
	Author      []string `json:"author"`
	PictureUrl  string   `json:"picture_url"`
	PictureID   int      `json:"picture_id"`
	PictureName string   `json:"picture_name"`
	VedioUrl    string   `json:"vedio_url"`
	VedioID     int      `json:"vedio_id"`
	VedioName   string   `json:"vedio_name"`
}

//AddCourseParam + CourseExtraParam
type AddCourseAllParam struct {
	Title     string `json:"title" binding:"required"`
	Desc      string `json:"desc"`
	AuthorID  int    `json:"author_id"`
	PictureID int    `json:"picture_id"`
	VedioID   int    `json:"vedio_id"`
}

//add course
type AddCourseParam struct {
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc"`
}

type UpdateCourseParam struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc"`
}

type CourseExtraParam struct {
	CourseID  int `json:"course_id"`
	AuthorID  int `json:"author_id"`
	PictureID int `json:"picture_id"`
	VedioID   int `json:"vedio_id"`
}

type AddCourseExtraRes struct {
	CourseResourceID   int   `json:"course_resource_id"`
	TeachersResourceID []int `json:"teachers_resource_id"`
	PictureResourceID  int   `json:"picture_resource_id"`
	VedioResourceID    int   `json:"vedio_resource_id"`
}

type AddCourseLessonParam struct {
	CourseID    int    `json:"course_id"`
	LessonTitle string `json:"lesson_title"`
	LessonDesc  string `json:"lesson_desc"`
}

type CourseLessonAllRes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type LessonItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type UpdateCourse struct {
	CourseID int `json:"course_id"`
	//basic
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc"`

	//extra
	AuthorID  int `json:"author_id"`
	PictureID int `json:"picture_id"`
	VedioID   int `json:"vedio_id"`
}
