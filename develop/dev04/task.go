package main

import (
	"sort"
	"strings"
)

type letters struct {
	word      string
	allLetter map[rune]int
}

func NewLetters(w string) letters {
	l := letters{}
	l.SetWord(w)
	return l
}

func (l *letters) Equal(EL letters) bool {
	for key, val := range l.allLetter {
		v, ok := EL.allLetter[key]
		if !ok {
			return false
		}
		if v != val {
			return false
		}
		continue
	}
	return true
}

func (l *letters) SetWord(w string) {
	l.word = w
	l.allLetter = make(map[rune]int)
	for _, v := range w {
		l.AddLetter(v)
	}
}

func (l *letters) AddLetter(char rune) {

	if _, ok := l.allLetter[char]; ok {
		l.allLetter[char]++
		return
	}
	l.allLetter[char] = 1

}

func SortMap(m map[string][]string) {
	for key, val := range m {
		if len(val) <= 1 {
			delete(m, key)
		}
		sort.Strings(val)
	}
}
func ToLowerCase(words []string) {
	for i, v := range words {
		words[i] = strings.ToLower(v)
	}
}

func SearchAnagramms(words []string) map[string][]string {
	wordsArr := make([]letters, 0)
	ToLowerCase(words)
	anagrammsArr := make(map[string][]string)
	addCheck := false
	for _, v := range words {
		lword := NewLetters(v)
		for j := 0; j < len(wordsArr); j++ {
			if wordsArr[j].word == "" {
				break
			}
			if lword.Equal(wordsArr[j]) {
				if lword.word != wordsArr[j].word {
					anagrammsArr[wordsArr[j].word] = append(anagrammsArr[wordsArr[j].word], v)
				}
				addCheck = !addCheck
				break
			}
		}
		if _, key := anagrammsArr[v]; !key && !addCheck {
			wordsArr = append(wordsArr, NewLetters(v))
			anagrammsArr[v] = []string{}
			continue
		}
		addCheck = !addCheck
	}
	SortMap(anagrammsArr)
	return anagrammsArr
}
