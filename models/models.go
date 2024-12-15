package models

type Course struct {
	Name    string // Course name
	Grade   string // Grade earned in the course (e.g., "A", "B", etc.)
	Credits int64  // Number of credits for the course
	Prerequisite string
	// QPTs    float64 // Quality Points for the course (Credits * Grade Points)
}
type Student struct {
	ID              string // Unique student ID
	SecondName      string // Student name
	FirstName       string // Student name
	Major           string
	StartYear       string
	GPA             float64
	CoursesTaken    map[string]Course // Map of course code to Course details
	RequiredCourses map[string]Course // Map of course code to Course details
}
