package dto

// add 接口入参数

// add 接口返回值
type AddCourseRes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// 课程推荐接口返回值
type CourseRecommendRes struct {
	PictureUrl string   `json:"picture_url"`
	CourseName string   `json:"course_name"`
	Teacher    []string `json:"teacher"`
	CourseDesc string   `json:"course_desc"`
	CourseID   int      `json:"course_id"`
}

// courseall 接口入参数
type CourseAllParam struct {
	CurrentPage int64 `json:"current_page" binding:"required"`
	PageSize    int64 `json:"page_size" binding:"required"`
}

// courseall 接口返回值
type CourseAllRes struct {
	TotalNumber int64 `json:"total_number"`
	// PageSize    int64        `json:"page_size"`
	// CurrentPage int64        `json:"current_page"`
	// Pages       []int64      `json:"pages"`
	CourseList []CourseRecommendRes `json:"course_list"`
}

/*
// 课程列表项
type CourseItem struct {
	CourseID   int    `json:"course_id"`
	Title      string `json:"course_name"`
	Desc       string `json:"course_desc"`
	Teacher       []string `json:"teacher"`
	PictureUrl string `json:"picture_url"`
}
*/
type Vedio struct {
	Url  string `json:"url"`
	Desc string `json:"desc"`
	Name string `json:"name"`
}

// 课程概览接口返回值
type CourseOverviewRes struct {
	Title      string       `json:"title"`
	Introduce  string       `json:"introduce"`
	Teacher    []string     `json:"teacher"`
	PictureUrl string       `json:"picture_url"`
	GuideVedio Vedio        `json:"guide_vedio"`
	LessonList []LessonItem `json:"lesson_list"`
}
