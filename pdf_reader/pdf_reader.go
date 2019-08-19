package pdf_reader

import (
	"github.com/cheggaaa/go-poppler"
	"log"
)

func Pdf2Text(filename string) string {
	var document, err = poppler.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	var numPages = document.GetNPages()

	var docText = ""

	for i := 0; i < numPages; i++ {
		var page = document.GetPage(i)
		docText += page.Text()
	}

	return docText
}