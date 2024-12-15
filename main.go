package main

import (
	"fmt"
	"os"

	comparator "audit/couseComparator"
	"audit/extractor"
	"audit/pdfparser" // Import the pdfparser package
	"audit/printer"

	"github.com/unidoc/unipdf/v3/common/license"
	// "github.com/pdfcpu/pdfcpu/pkg/api"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(`f443cd381251f3fc397d9bfd2574b78293aea9186f506eaa1bc9ab5b51b98ea5`)
	if err != nil {
		panic(err)
	}
}

func main() {

	// filename := "201853698_student_transcript"
	// filename := "202078151_student_transcript"
	filename := "201976068_student_transcript"

	pdfPath := "/Users/aigera/Downloads/" + filename + ".pdf" // Add ".pdf" extension if needed
	err, wholeTransr := pdfparser.ParsePdfText(pdfPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	//populate student struct with unique info
	student := extractor.ParseStudentInfo(wholeTransr)
	printer.PrintStudentInfo(student)
	// fmt.Println(student.FirstName)
	// fmt.Println(student.SecondName)
	// fmt.Println(student.ID)
	// fmt.Println(student.Major)
	// fmt.Println(student.StartYear)
	// fmt.Println(student.GPA)

	arrayOfCourseLines := extractor.ParseStudentCourses(wholeTransr)
	// fmt.Println(len(arrayOfCourseLines))
	// for _, line := range arrayOfCourseLines {
	// 	fmt.Printf("-------%s\n", line)
	// }
	// fmt.Println(arrayOfGrdLines)

	// fmt.Println("TAKEN from MAP: ")
	takenCourses := extractor.ExtractExactPassedCourses(arrayOfCourseLines)
	student.CoursesTaken = takenCourses //
	// fmt.Println(takenCourses)
	// fmt.Println(len(takenCourses))
	// found := takenCourses["MATH161"]
	// fmt.Println(found)

	// still required to be taken
	missingCourses := comparator.CheckRequiredCoreCourses(student)
	// fmt.Println("MISSING REQUIRED COURSES:")
	// fmt.Println(missingCourses)

	electiveCourses := comparator.GetTakenElectiveCourses(student)
	// fmt.Println("TAKEN All ELECTIVE COURSES:")
	// fmt.Println(electiveCourses)

	// fmt.Println("ALL ELECTIVE COURSES REQUIRED: ")
	// fmt.Println(db.ElectiveCourses)
	leftElectiveCourses := comparator.GetLeftElectiveCourses(electiveCourses)
	// fmt.Println(leftElectiveCourses)

	fmt.Println("FINAL RESULT - LEFT TO TAKE: ")
	printer.PrintCourses(missingCourses)
	printer.PrintCateoriesOfCoursesLeft(leftElectiveCourses)
	printer.PrintEntireStatistics(missingCourses, leftElectiveCourses)
}
