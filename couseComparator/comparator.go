package comparator

import (
	"audit/db"
	"audit/models"
)

// this function goes through the required courses and check how many
// and which required courses are still missing.
func CheckRequiredCoreCourses(student models.Student) map[string]models.Course {
	missingCourses := make(map[string]models.Course)

	for code, course := range db.CoreCourses {
		// Check if the course code is missing in the student's courses
		if _, found := student.CoursesTaken[code]; !found {
			missingCourses[code] = course
		}
	}

	return missingCourses
}

// remove all required/core courses from the list of taken courses by the student
// get all courses which are not core courses and need to be checked further
func GetTakenElectiveCourses(student models.Student) map[string]models.Course {
	takenNotCoreCourses := make(map[string]models.Course)

	for code, course := range student.CoursesTaken {
		// Check if the course code is missing in the student's courses
		if _, found := db.CoreCourses[code]; !found {
			takenNotCoreCourses[code] = course
		}
	}

	return takenNotCoreCourses
}

func GetLeftElectiveCourses(electiveCourses map[string]models.Course) map[string]int {
	for code, _ := range electiveCourses {
		if len(code) > 3 {
			coursePrefix := code[:len(code)-3]
			if coursePrefix == "CSCI" {
				if db.ElectiveCourses["Technical"] > 0 {
					db.ElectiveCourses["Technical"]--
				} else {
					db.ElectiveCourses["Open Elective"]--
				}

			}
			if _, found := db.TechnicalElectiveCourses[code]; found {
				if db.ElectiveCourses["Technical"] > 0 {
					db.ElectiveCourses["Technical"]--
				} else {
					db.ElectiveCourses["Open Elective"]--
				}
			}

			//communication core
			if _, found := db.CommunicationCoreCourses[code]; found {
				if db.ElectiveCourses["Communication Core"] > 0 {
					db.ElectiveCourses["Communication Core"]--
				} else {
					db.ElectiveCourses["Open Elective"]--
				}
			}

			// kazakh languages
			if coursePrefix == "KAZ" {
				if db.ElectiveCourses["Kazakh Language"] > 0 {
					db.ElectiveCourses["Kazakh Language"]--
				} else {
					db.ElectiveCourses["Open Elective"]--
				}
			}

			//natural electives
			if coursePrefix == "PHYS" || coursePrefix == "CHEM" || coursePrefix == "BIOL" || coursePrefix == "GEOL" {
				if db.ElectiveCourses["Natural Science"] > 0 {
					db.ElectiveCourses["Natural Science"]--
				} else {
					db.ElectiveCourses["Open Elective"]--
				}
			}

			if coursePrefix == "SOC" || coursePrefix == "PLS" || coursePrefix == "ANT" || coursePrefix == "ECON" || coursePrefix == "LING" {
				if db.ElectiveCourses["Social Science"] > 0 {
					db.ElectiveCourses["Social Science"]--
				} else {
					db.ElectiveCourses["Open Elective"]--
				}
			}
		}
	}
	return db.ElectiveCourses
}

func ComputeCreditsCourses(courses map[string]models.Course) (int, int) {
	var totCredits = 0

	for _, courses := range courses {
		totCredits += int(courses.Credits)
	}
	return len(courses), totCredits
}

func ComputeCreditsBasedCategories(categories map[string]int) (int, int) {
	var totCourses = 0
	var totCredits = 0
	for _, amount := range categories {
		totCourses += amount
		totCredits += amount * 6
	}
	return totCourses, totCredits
}
