// description: 学生
//
// author: vignetting
// time: 2021/5/10

package models

type Student struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name"`
	// 专业
	Major string `json:"major"`
	// 年级
	Grade uint `json:"grade"`
	Age   uint `json:"age"`
}

func GetStudent(id int) (Student, error) {
	var student Student
	return student, db.Take(&student, id).Error
}

func GetStudents(pageNumber, pageSize int) []Student {
	var students []Student
	db.Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&students)
	return students
}

func UpdateStudent(student Student) {
	db.Model(&student).Updates(student)
}
