package printer

import (
	"audit/comporator"
	"audit/models"
	"fmt"
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
