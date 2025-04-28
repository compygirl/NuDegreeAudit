package printer

import (
	"audit/comporator"
	"audit/models"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func PrintStudentInfo(student models.Student) {
	// Print student details with formatted alignment
	fmt.Println("\n=====================  STUDENT INFO  =====================")
	fmt.Printf("%-15s : %s\n", "First Name", student.FirstName)
	fmt.Printf("%-15s : %s\n", "Second Name", student.SecondName)
	fmt.Printf("%-15s : %s\n", "ID", student.ID)
	fmt.Printf("%-15s : %s\n", "Major", student.Major)
	fmt.Printf("%-15s : %s\n", "Start Year", student.StartYear)
	fmt.Printf("%-15s : %.2f\n", "GPA", student.GPA)
	fmt.Println(strings.Repeat("-", 50)) // Bottom border
}

func PrintCourses(courses map[string]models.Course) {
	fmt.Println("\n------------------------  COURSES  ------------------------")
	fmt.Printf("%-10s | %-40s | %6s\n", "Code", "Course Name", "Credits") // Header
	fmt.Println(strings.Repeat("-", 60))                                  // Separator

	for code, course := range courses {
		fmt.Printf("%-10s | %-40s | %6d\n", code, course.Name, course.Credits)
	}

	fmt.Println(strings.Repeat("-", 60)) // Footer Separator
	amountCourses, amountCredits := comporator.ComputeCreditsCourses(courses)
	PrintStatistics(amountCourses, amountCredits)
}

func PrintCateoriesOfCoursesLeft(categories map[string]int) {
	fmt.Println("\n------------------------  CATEGORIES  ------------------------")
	fmt.Printf("%-25s | %10s\n", "Category", "Amount") // Header
	fmt.Println(strings.Repeat("-", 40))               // Separator

	for courseType, amount := range categories {
		fmt.Printf("%-25s | %10d\n", courseType, amount)
	}

	fmt.Println(strings.Repeat("-", 40)) // Footer Separator
	amountCourses, amountCredits := comporator.ComputeCreditsBasedCategories(categories)
	PrintStatistics(amountCourses, amountCredits)
}

func PrintStatistics(amountCourses int, amountCredits int) {
	fmt.Printf(">> Remaining: %-5d COURSES | %-5d CREDITS\n\n", amountCourses, amountCredits)
}

func PrintEntireStatistics(courses map[string]models.Course, categories map[string]int) {
	amountCourses1, amountCredits1 := comporator.ComputeCreditsCourses(courses)
	amountCourses2, amountCredits2 := comporator.ComputeCreditsBasedCategories(categories)

	fmt.Println("\n------------------------  TOTAL STATISTICS  ------------------------")
	fmt.Printf(">> Total Courses: %-5d | Total Credits: %-5d\n",
		amountCourses1+amountCourses2, amountCredits1+amountCredits2)
	fmt.Println(strings.Repeat("=", 50))
}

func WriteToCSV(student models.Student, filename string, courses map[string]models.Course, categories map[string]int) {
	var csvFilename string
	if strings.HasSuffix(filename, ".pdf") {
		csvFilename = strings.TrimSuffix(filename, ".pdf") + ".csv"
		// fmt.Println(csvFilename) // Output: report.csv
	} else {
		fmt.Println("No .pdf suffix found")
	}
	csvPath := "/Users/aigera/Downloads/csvtranscripts/" + csvFilename

	file, err := os.Create(csvPath)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV Header
	header := []string{"First Name", "Second Name", "ID", "Major", "Start Year", "GPA"}
	writer.Write(header)

	// Write Student Rows
	row := []string{
		student.FirstName,
		student.SecondName,
		student.ID,
		student.Major,
		student.StartYear,
		fmt.Sprintf("%.2f", student.GPA),
	}
	writer.Write(row)

	header2 := []string{"Course Code", "Course Name", "Course Credits", "Grade"}
	writer.Write(header2)
	//add all courses in the format as COURSE_CODE COURSE_NAME CREDITS GRADE
	for code, course := range student.CoursesTaken {
		// fmt.Println("WRITER!!!!!")
		row := []string{
			code,
			course.Name,
			fmt.Sprintf("%d", course.Credits),
			course.Grade,
		}
		writer.Write(row)
	}
	header3 := []string{"MISSING REQUIRED COURSES:"}
	writer.Write(header3)

	header4 := []string{"Course Code", "Course Name", "Course Credits", "Grade"}
	writer.Write(header4)
	for code, course := range courses {
		// fmt.Println("WRITER!!!!!")
		row := []string{
			code,
			course.Name,
			fmt.Sprintf("%d", course.Credits),
			"N/A",
		}
		writer.Write(row)
	}

	header5 := []string{"MISSING CATEGORIES:"}
	writer.Write(header5)

	header6 := []string{"Category Name", "Amount"}
	writer.Write(header6)

	for courseType, amount := range categories {
		row := []string{
			courseType,
			fmt.Sprintf("%d", amount),
		}
		writer.Write(row)
	}

	//statistics :
	amountCourses1, amountCredits1 := comporator.ComputeCreditsCourses(courses)
	amountCourses2, amountCredits2 := comporator.ComputeCreditsBasedCategories(categories)
	header7 := []string{"STATISTICS: "}
	writer.Write(header7)

	final_row := []string{"TOTAL COURSES AMOUNT: ", fmt.Sprintf("%d", amountCourses1+amountCourses2), "TOTAL CREDITS: ", fmt.Sprintf("%d", amountCredits1+amountCredits2)}
	writer.Write(final_row)

}
