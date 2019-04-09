package tokens

import (
	"sort"
	"strings"
)

type TokensChecker struct {
	dictionary []string
}

func NewTokenChecker(wordList []string) (result *TokensChecker) {
	result = &TokensChecker{
		dictionary:wordList,
	}
	sort.Sort(sort.Reverse(sort.StringSlice(result.dictionary)))
	return result
}

func(t *TokensChecker) Dictionary() []string {
	return t.dictionary
}

// Cut repeatedly reversing to one step on fail
func(t *TokensChecker) Check(phrase string) bool {
	var (
		remain string
		indexes []int
	)
	remain = phrase
	for i, _ := range t.dictionary {
		remain, indexes = t.cut(remain, i)
		if len(remain) == 0 {
			return true
		}
		if len(indexes) > 0 {
			for len(indexes) > 0 {
				j := len(indexes) - 1
				remain = t.dictionary[indexes[j]] + remain
				if indexes[j] < len(t.dictionary) - 1 {
					remain, indexes = t.cut(remain, indexes[j] + 1)
					if len(remain) == 0 {
						return true
					}

				}
			}
		}
	}
	return false
}

// Cut until it's possible
// Return remaining string and indexes of cut tokens
func(t *TokensChecker) cut(remain string, startIndex int) (result string, indexes []int)  {
	var i int
	i = startIndex
	for i < len(t.dictionary) {
		if !strings.HasPrefix(remain, t.dictionary[i]) {
			i++
			continue
		}
		remain = remain[len(t.dictionary[i]):]
		indexes = append(indexes, i)
		i = 0
	}
	return remain, indexes
}