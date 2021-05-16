// description: 记录
//
// author: vignetting
// time: 2021/5/10

package models

import (
	"time"
)

type Record struct {
	ID int `json:"id"`

	StudentId int `json:"studentId" binding:"required"`

	CourseId int `json:"courseId"`

	Time time.Time `json:"time"`
}

// description: 获取记录
// param: studentId 学生id，小于零则不使用
// param: courseId 课程id，小于零则不使用
// param: date 时间，为 nil 则不使用
// param: pageNumber
// param: pageSize
// return: []Record
func GetRecords(studentId, courseId *int, date *time.Time, pageNumber, pageSize int) []Record {
	var records []Record
	tx := db.Limit(pageSize).Offset((pageNumber - 1) * pageSize)

	if studentId != nil {
		tx.Where("student_id = ?", *studentId)
	}

	if courseId != nil {
		tx.Where("course_id = ?", *courseId)
	}

	if date != nil {
		tx.Where("time >= ? and time <= ?", date, date.AddDate(0, 0, 1))
	}

	tx.Find(&records)

	return records
}

func InsertRecords(record Record) {
	db.Create(&record)
}
