package main

import (
	"audit/comporator"
	"audit/extractor"
	"audit/pdfparser"
	"audit/printer"
	"fmt"
	"os"
	"path/filepath"

	"github.com/unidoc/unipdf/v3/common/license"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	pdf_extractor_key := `738430bbb8e450415052758235b3cc9a5f2453b1a55bed14faa9357efc1090ec`
	err := license.SetMeteredKey(pdf_extractor_key)
	if err != nil {
		panic(err)
	}
}

func main() {

	pdfDir := "/Users/aigera/Downloads/transcripts"
	// Read all files in the directory
	files, err := os.ReadDir(pdfDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	// Loop through all files
	for _, file := range files {
		// Check if the file has a .pdf extension
		if filepath.Ext(file.Name()) == ".pdf" {
			pdfPath := filepath.Join(pdfDir, file.Name()) // Full path to the PDF file

			// Parse the PDF content
			err, wholeTransr := pdfparser.ParsePdfText(pdfPath)
			if err != nil {
				fmt.Printf("Error parsing %s: %v\n", file.Name(), err)
				continue // Skip to the next file on error
			}

			// Parse student info
			student := extractor.ParseStudentInfo(wholeTransr)

			// Print student information
			fmt.Println("Processing file:", file.Name())
			arrayOfCourseLines := extractor.ParseStudentCourses(wholeTransr)
			student.CoursesTaken = extractor.ExtractExactPassedCourses(arrayOfCourseLines)
			fmt.Println(student.CoursesTaken)

			// retrieve courses that still required to be taken by student
			missingCourses := comporator.CheckRequiredCoreCourses(student)
			electiveCourses := comporator.GetTakenElectiveCourses(student)
			leftElectiveCourses := comporator.GetLeftElectiveCourses(electiveCourses)

			printer.PrintStudentInfo(student)
			printer.PrintCourses(missingCourses)
			printer.PrintCateoriesOfCoursesLeft(leftElectiveCourses)
			printer.PrintEntireStatistics(missingCourses, leftElectiveCourses)
		}
	}
}
