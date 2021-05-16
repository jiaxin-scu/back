// description:
//
// author: vignetting
// time: 2021/5/10

package models

import (
	"structure/pkg/e"
	"time"
)

type Course struct {
	ID          int `json:"id" binding:"required"`
	TeacherId   int `json:"teacherId"`
	ClassRoomId int `json:"classRoomId"`
	StartNumber int `json:"startNumber"`
	EndNumber   int `json:"endNumber"`
	Week        int `json:"week"`
}

func getNumber() int {
	minutes := time.Now().Hour()*60 + time.Now().Minute()

	switch {
	case minutes > 490 && minutes < 545:
		return 1
	case minutes > 545 && minutes < 600:
		return 2
	case minutes > 610 && minutes < 665:
		return 3
	case minutes > 665 && minutes < 720:
		return 4
	case minutes > 825 && minutes < 880:
		return 5
	case minutes > 880 && minutes < 935:
		return 6
	case minutes > 945 && minutes < 1000:
		return 7
	case minutes > 1000 && minutes < 1055:
		return 8
	case minutes > 1055 && minutes < 1110:
		return 9
	case minutes > 1155 && minutes < 1210:
		return 10
	case minutes > 1210 && minutes < 1265:
		return 11
	case minutes > 1265 && minutes < 1320:
		return 12
	}
	return 0
}

func GetCourseId(studentId, classRoomId int) (int, error) {
	var course Course
	number := getNumber()
	week := int(time.Now().Weekday())

	db.Raw("select id from course, course_student where class_room_id = ? and start_number <= ? and end_number >= ? and week = ? and id = course_id and student_id = ?", classRoomId, number, number, week, studentId).Scan(&course)

	if course.ID > 0 {
		return course.ID, nil
	}
	return 0, e.Fail("未查到相符合的课程")
}

func GetStudentCourses(studentId, pageNumber, pageSize int) []Course {
	var courses []Course
	db.Raw("select id, teacher_id, class_room_id, start_number, end_number, week from course, course_student where student_id = ? and course_id = id limit ?, ?", studentId, (pageNumber-1)*pageSize, pageSize).Scan(&courses)
	return courses
}

func GetCourses(teacherId, classRoomId, pageNumber, pageSize int) []Course {
	var courses []Course
	tx := db.Limit(pageSize).Offset((pageNumber - 1) * pageSize)
	if teacherId > 0 {
		tx.Where("teacher_id = ?", teacherId)
	}
	if classRoomId > 0 {
		tx.Where("class_room_id = ?", classRoomId)
	}
	tx.Find(&courses)
	return courses
}
