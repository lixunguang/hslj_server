package dto

type LessonIDParam struct {
	ID int `json:"id"`
}

type LessonParam struct {
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc"`
}

type LessonResourceAllRes struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Type    int    `json:"type"`
}

type LessonSectionContent struct {
	LessonID    int    `json:"lesson_id"`
	ResourceID  int    `json:"resource_id"`
	Desc        string `json:"desc"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentType int    `json:"content_type"`
	Index       int    `json:"content_index"` //用于展示时排序或者资源用途（当为作业时）。
}

type UpdateLessonParam struct {
	CourseID int    `json:"course_id"`
	LessonID int    `json:"lesson_id"`
	Title    string `json:"title" `
	Desc     string `json:"desc"`
}

type LessonResource struct {
	ResourceID    int `json:"resource_id"`
	ResourceType  int `json:"resource_type"` //可选的type为：1 教学内容，2 参考资源，3 实验资源，4作业
	ResourceIndex int `json:"resource_index"`
}

type LessonResourceParam struct {
	LessonID           int              `json:"lesson_id"`
	CourseID           int              `json:"course_id"`
	LessonResourceList []LessonResource `json:"lesson_resource_list"`
}

type LessonContentParam struct {
	CourseID   int `json:"course_id"`
	LessonID   int `json:"lesson_id"`
	ResourceID int `json:"resource_id"`
}

type LessonResourceAddParam struct {
	CourseID      int `json:"course_id"`
	LessonID      int `json:"lesson_id"`
	ResourceID    int `json:"resource_id"`
	ResourceIndex int `json:"resource_index"`
}

type LessonResourceDelParam struct {
	CourseID   int `json:"course_id"`
	LessonID   int `json:"lesson_id"`
	ResourceID int `json:"resource_id"`
}
