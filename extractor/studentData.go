// pdfparser/student_parser.go
package extractor

import (
	// import the models package
	"audit/db"
	"audit/models"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// // ParseStudentDetails extracts student information from the given text
// func ParseStudentDetails(text string) models.Student {
// 	var student models.Student

// 	// Extract the full name
// 	nameStart := "Student Name: "
// 	// nameEnd := "School: "
// 	fullName := extractStringBetween(text, nameStart, nameEnd)
// 	nameParts := strings.Split(fullName, " ")
// 	if len(nameParts) >= 2 {
// 		student.LastName = nameParts[0]
// 		student.FirstName = nameParts[1]
// 	}

// 	// Extract primary major
// 	majorStart := "Primary Major: "
// 	majorEnd := "Admission"
// 	student.PrimaryMajor = extractStringBetween(text, majorStart, majorEnd)

// 	// Extract ID
// 	idStart := "Program:\n"
// 	idEnd := "\nBachelor"
// 	student.ID = extractStringBetween(text, idStart, idEnd)

// 	return student
// }

// ParseStudentInfo parses the student information from a given text
func ParseStudentInfo(text string) models.Student {
	var student models.Student

	// Extract student name
	studentNamePart := extractAfter(text, "Student Name:")
	nameParts := strings.Fields(studentNamePart)

	if len(nameParts) >= 2 {
		student.SecondName = nameParts[0]
		student.FirstName = nameParts[1]
	}

	// Extract primary major
	primaryMajor := extractAfter(text, "Primary Major:")
	student.Major = strings.Split(primaryMajor, "\n")[0] // Get only the first line

	// Extract ID
	id := extractAfter(text, "Program:")
	student.ID = extractNumber(id) // Extract only the numeric part

	// Extract START YEAR
	startSemester := extractAfter(text, "semester:")
	semester := strings.Fields(startSemester)
	fmt.Println(semester[1])
	if len(semester) >= 2 {
		student.StartYear = semester[1]
	}

	// Extract GPA
	GPAstring := extractAfter(text, "Cumulative GPA:")
	arrayFORGPA := strings.Fields(GPAstring)
	if len(nameParts) >= 1 {
		GPAstring = arrayFORGPA[0]
	}
	GPAfloat, err := strconv.ParseFloat(GPAstring, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return student
	}
	student.GPA = GPAfloat

	return student

}

// Helper function to extract text after a specific keyword
func extractAfter(text, keyword string) string {
	index := strings.Index(text, keyword)
	if index == -1 {
		return ""
	}
	return strings.TrimSpace(text[index+len(keyword):])
}

// Helper function to extract numeric part from a string
func extractNumber(text string) string {
	re := regexp.MustCompile(`\d+`) // Match digits
	match := re.FindString(text)
	if match != "" {
		return match
	}
	return ""
}

func ParseStudentCourses(text string) []string {
	var subjGrades string
	subjGrades = extractAfter(text, "Bachelor of\nScience in\nComputer Science")
	lines := strings.Split(subjGrades, "\n")

	//combie into 2D array
	arrays := [][]string{db.PrefixesRequired, db.PrefixTechnicalElectives, db.PrefixNaturalElectives, db.PrefixKAZElectives, db.PrefixOpenElectives, db.PrefixSocialElectives}

	//combine all db.Prefixes into 1D array
	var combinedPrefixes []string
	for _, array := range arrays {
		combinedPrefixes = append(combinedPrefixes, array...)
	}

	//exytacting only course related lines
	courses := extractMatchingStrings(lines, combinedPrefixes)
	return courses
}

func extractMatchingStrings(lines []string, prefixes []string) []string {
	var matched []string

	for _, line := range lines {
		for _, prefix := range prefixes {
			if strings.HasPrefix(line, prefix) {
				matched = append(matched, line)
				break // No need to check other prefixes once a match is found
			}
		}
	}

	return matched
}

func ExtractExactPassedCourses(lines []string) map[string]models.Course {
	courses := make(map[string]models.Course)
	for _, line := range lines {
		//maps's key - course code
		takenCourseArr := strings.Fields(line)
		courseCode := takenCourseArr[0] + takenCourseArr[1]
		//the course struct
		var course models.Course
		length := len(takenCourseArr)

		// exception for the Internships
		if courseCode == "CSCI299" || courseCode == "CSCI399" {
			//extract course Letter Grade
			course.Grade = takenCourseArr[length-2]

			//extract credits for each course
			creditsStr := takenCourseArr[length-1]
			creditsInt, _ := strconv.ParseInt(creditsStr, 10, 64)
			course.Credits = creditsInt
		} else {
			//extract course Letter Grade
			course.Grade = takenCourseArr[length-3]

			//extract credits for each course
			creditsStr := takenCourseArr[length-2]
			creditsInt, _ := strconv.ParseInt(creditsStr, 10, 64)
			course.Credits = creditsInt
		}
		courses[courseCode] = course

	}
	return courses
}
