package transform

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/samirgadkari/search/pkg/data"
)

type ProcFunc struct {
	ToLower            func(string) string
	Replace            func(string) string
	GetWords           func(string, []string) []string
	WordsToInts        func([]string, []data.WordInt) []data.WordInt
	WriteWordInts      func(string, int, []data.WordInt)
	GetWordIntMappings func() (map[string]data.WordInt, map[data.WordInt]string)
	RemoveStopwords    func([]string) []string
}

const (
	maxWordsPerDoc    = 1024
	wordToIntFilename = "wordToInt.txt"
	intToWordFilename = "intToWord.txt"
)

func replacer(s string) string {
	// This character replacer replaces:
	// an apostrophe ' with empty string,
	// a underscore _ with space,
	// and all other punctuation marks with space.
	// We're also going to remove all numbers.
	r := strings.NewReplacer("'", "", "_", " ", "!", " ",
		"!", " ", "\"", " ", "#", " ", "$", " ",
		"%", " ", "&", " ",
		"(", " ", ")", " ", "*", " ", "+", " ",
		",", " ", "-", " ", ".", " ", "/", " ",
		":", " ", ";", " ", "<", " ", "=", " ",
		">", " ", "?", " ", "@", " ", "[", " ",
		"\\", " ", "]", " ", "^", " ",
		"_", " ", "`", " ", "{", " ", "|", " ",
		"}", " ", "~", " ",
		"0", "", "1", "", "2", "", "3", "", "4", "",
		"5", "", "6", "", "7", "", "8", "", "9", "")

	return r.Replace(s)
}

func getWordsFn() func(string, []string) []string {

	re := regexp.MustCompile(`(\w+)`)

	return func(s string, wordMatches []string) []string {
		matches := re.FindAllStringSubmatch(s, -1)
		if matches != nil {
			wordMatches = wordMatches[:0]
			for _, match := range matches {
				if len(match) != 2 {
					fmt.Printf("len(match) is %d which is invalid\n", len(match))
					os.Exit(-1)
				}
				wordMatches = append(wordMatches, match[1])
			}

			if len(wordMatches) > maxWordsPerDoc {
				fmt.Printf("Too many words in doc (%d) !!", len(wordMatches))
				os.Exit(-1)
			}

			return wordMatches
		} else {
			return nil
		}
	}
}

func removeStopwordsFn(stopwords []string) func([]string) []string {

	swMap := make(map[string]struct{}, len(stopwords))
	for _, v := range stopwords {
		swMap[v] = struct{}{}
	}

	return func(words []string) []string {

		var result []string
		for _, word := range words {
			_, ok := swMap[word]
			if !ok {
				result = append(result, word)
			}
		}

		return result
	}
}

func wordToIntsFns() (func([]string, []data.WordInt) []data.WordInt,
	func() (map[string]data.WordInt, map[data.WordInt]string)) {

	wordToInt := make(map[string]data.WordInt)
	intToWord := make(map[data.WordInt]string)
	var wordNum data.WordInt = 0
	return func(words []string, wordInts []data.WordInt) []data.WordInt {

			wordInts = wordInts[:0]
			for _, word := range words {
				if _, ok := wordToInt[word]; ok == false {
					wordToInt[word] = wordNum
					intToWord[wordNum] = word
					wordInts = append(wordInts, wordNum)
					wordNum += 1
				} else {
					wordInts = append(wordInts, wordToInt[word])
				}

			}

			return wordInts
		}, func() (map[string]data.WordInt, map[data.WordInt]string) {
			return wordToInt, intToWord
		}
}

func GenProcFunc(stopwords []string) *ProcFunc {

	var procFunc ProcFunc

	procFunc.Replace = replacer

	procFunc.ToLower = func(s string) string {
		return strings.ToLower(s)
	}

	procFunc.GetWords = getWordsFn()
	procFunc.RemoveStopwords = removeStopwordsFn(stopwords)
	procFunc.WordsToInts, procFunc.GetWordIntMappings = wordToIntsFns()

	return &procFunc
}

func WordsToInts(loadStopwords func() []string,
	loadData func() (*data.Doc, bool),
	writeWordIntMappings func(map[string]data.WordInt, map[data.WordInt]string),
	storeData func(*data.Doc, []data.WordInt),
	closeData func()) {

	proc := GenProcFunc(loadStopwords())

	words := make([]string, maxWordsPerDoc)
	wordInts := make([]data.WordInt, maxWordsPerDoc)
	var line string

	for {
		v, ok := loadData()
		if !ok {
			break
		}
		line = v.Text

		line = proc.Replace(line)
		line = proc.ToLower(line)
		words = proc.GetWords(line, words)
		words = proc.RemoveStopwords(words)
		wordInts = proc.WordsToInts(words, wordInts)

		storeData(v, wordInts)
	}

	closeData()
	wordToInt, intToWord := proc.GetWordIntMappings()
	writeWordIntMappings(wordToInt, intToWord)
}
