package cli

import (
	"fmt"
	"github.com/rishisubramanian/PDF-Summary/pdf-reader"
	"github.com/rishisubramanian/PDF-Summary/summarizer"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app                  = kingpin.New("PDF-Summary", "A command line app to summarize PDF articles.")
	filename     *string = kingpin.Flag("filename", "Name of the file to process.").Required().ExistingFile()
	numSentences *int    = kingpin.Flag("n", "Number of sentences to return.").Default("5").Int()
)

func main() {
	kingpin.Version("0.0")
	kingpin.Parse()

	pdfText := pdf_reader.Pdf2Text(*filename)

	summaryLines := summarizer.SummarizeText(pdfText, *numSentences)

	fmt.Println()
	for index, sentence := range summaryLines {
		fmt.Println(index, sentence)
	}
}