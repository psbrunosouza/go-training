package challenges

import "fmt"

var dnaStrings map[string]string = map[string]string{
	"T": "A",
	"A": "T",
	"G": "C",
	"C": "G",
}

func DNAStrand(dna string) (result string) {
	for _, v := range dna {
		result += dnaStrings[fmt.Sprintf("%c", v)]
	}
	return
}