package dto

// lesson参数
type Lesson struct {
	Title string `json::"title"`
	Desc  string `json:"desc"`
	// 	Index    int    `json:"index"`
	// 	CourseID int    `json:"course_id"`
}

// 课程包含多堂课，每个课堂用LessonItem表示
type LessonItem struct {
	LessonTitle string `json:"lesson_title"`
	Desc        string `json:"desc"`
	LessonId    int    `json:"lesson_id"`
	Index       int    `json:"index"`
}

// 课堂概览接口返回值
type LessonOverviewRes struct {
	LessonTitle    string          `json:"lesson_title"`
	LessonSections []LessonSection `json:"lesson_sections"`
}
