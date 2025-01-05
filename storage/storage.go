package storage

import "lmsmodule/models"

var Courses = []models.Course{
	{ID: 1, Title: "Go Basics", Description: "Learn the basics of Go programming."},
	{ID: 2, Title: "Advanced Go", Description: "Deep dive into Go."},
}

var Assignments = []models.Assignment{
	{ID: 1, CourseID: 1, Title: "Introduction to Go"},
	{ID: 2, CourseID: 1, Title: "Variables and Types"},
}

var UserProgress = map[int]*models.UserProgress{
	1: {UserID: 1, Completed: map[int]bool{}, LastActivity: ""},
}
