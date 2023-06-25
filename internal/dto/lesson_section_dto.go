package dto

//课堂内容包含几个部分：如教学内容，教学参考资源，实验资源，作业
type LessonSection struct {
	//CourseID int `json:"course_id"` //新增0406
	LessonID int    `json:"lesson_id"`
	Title    string `json:"title"`
	Index    int    `json:"index"`
}

//包括内容，做实验，参考
type LessonSectionContent struct {
	LessonID    int    `json:"lesson_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentType int    `json:"content_type"`
	Index       int    `json:"index"` //用于展示时排序或者资源用途（当为作业时）。
}

type WorkContentItem struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentType int    `json:"content_type"` //在common define里面有定义：1图片，	2 word文档，	3 pdf文档，	4 apps文件，	5 ibe文件，	6文本（未用）	7 富文本。	8 视频
}

//作业重新定义一个结构体
type LessonSectionContentWork struct {
	LessonID    int               `json:"lesson_id"`
	Requirement string            `json:"requirement"`
	Contents    []WorkContentItem `json:"content"`
}
