package main

import (
	"flag"
	"tasks/tokens"
	"strings"
)

type Conf struct {
	phrase string
	wordListString string
	wordList []string
}

var conf Conf

func init()  {
	flag.StringVar(&conf.wordListString, `list`, `a what an nice day`, `Please pass a list of tokens divided by spaces`)
	flag.StringVar(&conf.phrase, `phrase`, `whataniceday`, `Please pass a phrase`)
	flag.Parse()
}

func main()  {
	var checker *tokens.TokensChecker
	conf.wordList = strings.Split(conf.wordListString, ` `)
	checker = tokens.NewTokenChecker(conf.wordList)
	println(checker.Check(conf.phrase))
}
