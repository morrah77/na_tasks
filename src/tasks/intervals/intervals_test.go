package intervals_test

import (
	"testing"
	"gotest.tools/assert"
	"tasks/intervals"
)

func TestNewIntervalMergerCreateSuccess(t *testing.T) {
	var (
		sourceIntervals = [][]int{
		{1, 2},
		{3, 4},
		{7, 12},
		{9, 14},
		{8, 15},
		{20,25},
		{18,30},
	}
	expectedIntervals = [][]int{
		{1, 2},
		{3, 4},
		{7, 12},
		{8, 15},
		{9, 14},
		{18,30},
		{20,25},
	}
	im *intervals.IntervalMerger
	err error
	)
	im, err = intervals.NewIntervalMerger(sourceIntervals);
	assert.NilError(t, err, `Should create IntervalMerger correctly`)
	assert.DeepEqual(t, im.Intervals(), expectedIntervals)
}

func TestNewIntervalMergerCreateErrorEmptySource(t *testing.T) {
	var (
		sourceIntervals = [][]int{{}}
		err error
	)
	_, err = intervals.NewIntervalMerger(sourceIntervals);
	assert.Error(t, err, `intervals: Invalid interval: []`, `Should fail with error on IntervalMerger creation`)
}

func TestNewIntervalMergerCreateErrorTooLongSource(t *testing.T) {
	var (
		sourceIntervals = [][]int{
			{1, 2, 8},
			{3, 4},
			{7, 12},
			{9, 14},
			{8, 15},
			{20,25},
			{18,30},
		}
		err error
	)
	_, err = intervals.NewIntervalMerger(sourceIntervals);
	assert.Error(t, err, `intervals: Invalid interval: [1,2,8]`, `Should fail with error on IntervalMerger creation`)
}

func TestNewIntervalMergerCreateErrorToShortSource(t *testing.T) {
	var (
		sourceIntervals = [][]int{
			{1, 2},
			{3},
			{7, 12},
		}
		err error
	)
	_, err = intervals.NewIntervalMerger(sourceIntervals);
	assert.Error(t, err, `intervals: Invalid interval: [3]`, `Should fail with error on IntervalMerger creation`)
}

func TestNewIntervalMergerMerge(t *testing.T) {
	var (
		sourceIntervals = [][]int{
			{1, 2},
			{3, 4},
			{7, 12},
			{9, 14},
			{8, 15},
			{20,25},
			{18,30},
			{35, 40},
			{37, 38},
			{42,43},
			{45, 46},
		}
		expectedIntervals = [][]int{
			{1, 4},
			{7, 15},
			{18,30},
			{35, 40},
			{42,43},
			{45, 46},
		}
		im *intervals.IntervalMerger
		err error
	)
	im, err = intervals.NewIntervalMerger(sourceIntervals);
	assert.NilError(t, err, `Should create IntervalMerger correctly`)
	im.Merge()
	assert.DeepEqual(t, im.Intervals(), expectedIntervals)
}