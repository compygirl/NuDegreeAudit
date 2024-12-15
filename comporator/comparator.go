package comporator

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
	dbElectiveCourses := make(map[string]int)
	for k, v := range db.ElectiveCourses {
		dbElectiveCourses[k] = v
	}
	for code, _ := range electiveCourses {
		if len(code) > 3 {
			coursePrefix := code[:len(code)-3]
			if coursePrefix == "CSCI" {
				if dbElectiveCourses["Technical"] > 0 {
					dbElectiveCourses["Technical"]--
				} else {

					dbElectiveCourses["Open Elective"]--
				}

			}
			if _, found := db.TechnicalElectiveCourses[code]; found {
				if dbElectiveCourses["Technical"] > 0 {
					dbElectiveCourses["Technical"]--
				} else {

					dbElectiveCourses["Open Elective"]--
				}
			}

			//communication core
			if _, found := db.CommunicationCoreCourses[code]; found {
				if dbElectiveCourses["Communication Core"] > 0 {
					dbElectiveCourses["Communication Core"]--
				} else {

					dbElectiveCourses["Open Elective"]--
				}
			}

			// kazakh languages
			if coursePrefix == "KAZ" {
				if dbElectiveCourses["Kazakh Language"] > 0 {
					dbElectiveCourses["Kazakh Language"]--
				} else {

					dbElectiveCourses["Open Elective"]--
				}
			}

			//natural electives
			if coursePrefix == "PHYS" || coursePrefix == "CHEM" || coursePrefix == "BIOL" || coursePrefix == "GEOL" {
				if dbElectiveCourses["Natural Science"] > 0 {
					dbElectiveCourses["Natural Science"]--
				} else {

					dbElectiveCourses["Open Elective"]--
				}
			}

			if coursePrefix == "SOC" || coursePrefix == "PLS" || coursePrefix == "ANT" || coursePrefix == "ECON" || coursePrefix == "LING" {
				if dbElectiveCourses["Social Science"] > 0 {
					dbElectiveCourses["Social Science"]--
				} else {

					dbElectiveCourses["Open Elective"]--
				}
			}
		}
	}
	return dbElectiveCourses
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
