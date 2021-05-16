// description: 老师
//
// author: vignetting
// time: 2021/5/10

package models

type Teacher struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name"`
	// 专业
	Major string `json:"major"`
	// 年级
	Grade uint `json:"grade"`
	// 性别，false 表示男，true 表示女
	Gender bool `json:"gender"`
}

func GetTeacher(id int) (Teacher, error) {
	var teacher Teacher
	return teacher, db.Take(&teacher, id).Error
}

func GetTeachers(pageNumber, pageSize int) []Teacher {
	var teachers []Teacher
	db.Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&teachers)
	return teachers
}

func UpdateTeacher(teacher Teacher) {
	db.Model(&teacher).Updates(teacher)
}
