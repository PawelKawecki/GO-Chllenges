package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"sort"
)

const wordsUrl = "https://pastebin.com/raw/jSD873gL"

var wordsSlice = []string {"mkeart", "sleewa", "edcudls", "iragoge", "usrlsle", "nalraoci", "nsdeuto", "amrhat", "inknsy", "iferkna"}


func main()  {
	fmt.Println("hello challenge #3 hard")

	words := getExtractedWords()

	found := searchWords(words)

	fmt.Println(found, len(found))

}

func searchWords(words []string) (found []string) {

	found = make([]string, len(wordsSlice))
	index := 0

	for _, word := range words {
		for _, val := range wordsSlice {
			if len(val) != len(word) {
				continue
			}

			if sortString(val) != sortString(word) {
				continue
			}

			found[index] = word

			index++
		}
	}

	return
}

func getExtractedWords() []string {
	resp, err := http.Get(wordsUrl)

	if err != nil {
		fmt.Println("There was a problem with connection to given url")
	}

	defer resp.Body.Close()

	words, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("There was a problem with extracting words")
	}

	return strings.Fields(string(words))
}

func sortString(w string) string {
	s := strings.Split(w, "")

	sort.Strings(s)

	return strings.Join(s, "")
}
