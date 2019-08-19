package summarizer

import (
	"gopkg.in/jdkato/prose.v2"
	"gonum.org/v1/gonum/mat"

)

func SummarizeText(text string, numLines int) []string {

	sentenceList := sentenceTokenizer(text)

	const lexrankThreshold float64 = 0.3
	ranks := lexrank(sentenceList, lexrankThreshold)

	topNIndices := topN(ranks, numLines)

	topNSentences := make([]string, numLines)

	for rankIdx, sentIdx := range topNIndices {
		topNSentences[rankIdx] = sentenceList[sentIdx]
	}

	return topNSentences
}

func sentenceTokenizer(text string) []string {
	doc, _ := prose.NewDocument(text)
	sentenceTexts := make([]string, len(text))

	for idx, sent := range doc.Sentences() {
		sentenceTexts[idx] = sent.Text
	}
	return sentenceTexts
}

func cosineIdf(sentence1 string, sentence2 string) float64 {

}

func powerMethod(distances [][]float64) []float64 {
	const errorTolerance float64 = 10 ^ -8

	length := len(distances)

	evecPrev := mat.NewVecDense(length, nil)
	evecPrev.SetVec()
	evecCurr := mat.NewVecDense(length, nil)

	for {

	}
}

func lexrank(sentenceList []string, threshold float64) []float64 {
	length := len(sentenceList)
	distanceMatrix := make([][]float64, length)
	for i := 0; i < length; i++ {
		distanceMatrix[i] = make([]float64, length)
	}

	//distanceMatrix := mat.NewDense(length, length, nil)

	degree := make([]float64, length)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			distanceMatrix[i][j] = cosineIdf(sentenceList[i], sentenceList[j])
			if distanceMatrix[i][j] > threshold {
				distanceMatrix[i][j] = 1
				degree[i]++
			} else {
				distanceMatrix[i][j] = 0
			}
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			distanceMatrix[i][j] = distanceMatrix[i][j] / degree[i]
		}
	}

	return powerMethod(distanceMatrix)
}

func topN(weights []float64, numResults int) []int {
	indexBuffer := make([]int, numResults)


}
