package main

import (
	"flag"
	"tasks/helpers"
	"tasks/intervals"
	"fmt"
)

var (
	listString string
	list [][]int
)

func init()  {
	flag.StringVar(&listString, `list`, `[[2,6],[8,10],[1,3],[15,18],[18,21]]`, `Please pass a list of intervals`)
	flag.Parse()
}

func main()  {
	var err error
	list, err = helpers.ParseList(listString)
	failOnError(err)
	im, err := intervals.NewIntervalMerger(list)
	failOnError(err)
	err = im.Merge()
	failOnError(err)
	fmt.Printf("%v\n", im.Intervals())
}

func failOnError(err error)  {
	if err != nil {
		panic(err)
	}
}

