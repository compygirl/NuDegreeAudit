package pdfparser

import (
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

// outputPdfText prints out contents of PDF file to stdout.
func ParsePdfText(inputPath string) (error, string) {
	f, err := os.Open(inputPath)
	if err != nil {
		return err, ""
	}

	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return err, ""
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err, ""
	}

	// fmt.Printf("--------------------\n")
	// fmt.Printf("PDF to text extraction:\n")
	// fmt.Printf("--------------------\n")
	var entireTransctipt string
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err, ""
		}

		ex, err := extractor.New(page)
		if err != nil {
			return err, ""
		}

		text, err := ex.ExtractText()
		if err != nil {
			return err, ""
		}
		entireTransctipt += text

		// fmt.Println("------------------------------")
		// fmt.Printf("Page %d:\n", pageNum)
		// fmt.Printf("\"%s\"\n", text)
		// fmt.Println("------------------------------")
	}

	return nil, entireTransctipt
}
