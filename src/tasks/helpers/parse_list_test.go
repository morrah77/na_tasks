package helpers_test

import (
"testing"
"gotest.tools/assert"
"tasks/helpers"
)

func TestParseListSuccess(t *testing.T) {
	var (
		stringList string = `[[2,6],[8,10],[1,3],[15,18],[18,21]]`
		expectedList [][]int = [][]int{
			{2, 6},
			{8, 10},
			{1, 3},
			{15, 18},
			{18, 21},
		}
		result [][]int
		err error
	)
	result, err = helpers.ParseList(stringList)
	assert.NilError(t, err, `Should parse list cuccessfully`)
	assert.DeepEqual(t, result, expectedList)
}

func TestParseListFailNotJson(t *testing.T) {
	var (
		stringList string = `[2,6],[8,10],[1,3],[15,18],[18,21]`
	)
	_, err := helpers.ParseList(stringList)
	assert.Error(t, err, `helpers: incorrect input`, `Should not parse list cuccessfully`)
}

func TestParseListFailNotCorrectNumSlice(t *testing.T) {
	var (
		stringList string = `[[2.5,6],[8,10],[1,3],[15,18],[18,21]]`
	)
	_, err := helpers.ParseList(stringList)
	assert.Error(t, err, `helpers: incorrect input`, `Should not parse list cuccessfully`)
}

func TestParseListFailNotCorrectSlice(t *testing.T) {
	var (
		stringList string = `[["a","b"],["c","d"]]`
	)
	_, err := helpers.ParseList(stringList)
	assert.Error(t, err, `helpers: incorrect input`, `Should not parse list cuccessfully`)
}
