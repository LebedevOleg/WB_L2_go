package main

type letters struct {
	allLetter map[rune]int
}

func (l letters) CheckLetters(word string) bool {
	for _, v := range word {
		if count, ok := l.allLetter[v]; ok {
			if count <= 0 {
				return false
			}
			l.allLetter[v]--
			if l.allLetter[v] == 0 {
				delete(l.allLetter, v)
			}
			continue
		}
		return false
	}
	return true
}

func (l *letters) AddLetter(char rune) {
	l.allLetter[char]++
	/* if _, ok := l.allLetter[char]; ok {
		l.allLetter[char]++
		return
	}
	l.allLetter[char] = 1 */

}

func SearchAnagramms(words []string) map[string][]string {

}
