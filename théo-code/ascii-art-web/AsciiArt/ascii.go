package ascii

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Ascii struct {
	Input string
	Font  string
}

func CreateOutput(output, file []byte, word string, m int) []byte {
	if m == 8 {
		return output
	}

	for j, n := 0, len(word); j < n; j++ {
		numOfNl := 0
		a := int(word[j]-32)*9 + 2 + m
		for i, l := 0, len(file); i < l; i++ {
			if file[i] == 13 {
				continue
			}
			if file[i] == 10 {
				numOfNl++
			} else if numOfNl == a-1 {
				output = append(output, file[i])
			} else if numOfNl == a {
				break
			}
		}
	}
	output = append(output, '\n')
	return CreateOutput(output, file, word, m+1)
}

func AsciiOutput(input, font string) (string, int) {
	var (
		wordsArr []string
		output   []byte
		index    int
	)
	file, err := ioutil.ReadFile("./assets/banners/" + font + ".txt")
	if err != nil {
		return "", http.StatusInternalServerError
	}
	wordsArr = strings.Split(input, "\n")
	for _, word := range wordsArr {
		output = CreateOutput(output, file, word, index)
	}
	return string(output), http.StatusOK
}
