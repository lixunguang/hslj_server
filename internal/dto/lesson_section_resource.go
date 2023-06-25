package dto

// LessonSectionResource参数
type LessonSectionResource struct {
	SectionID  int `json:"lesson_section_id"`
	ResourceID int `json:"resource_id"`
	Index      int `json:"index"`
}

type LessonSectionDelResource struct {
	SectionID  int `json:"lesson_section_id"`
	ResourceID int `json:"resource_id"`
}

type LessonSectionWork struct {
	SectionID  int `json:"lesson_section_id"`
	ResourceID int `json:"resource_id"`
}
