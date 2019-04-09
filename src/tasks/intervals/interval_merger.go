package intervals

import (
	"errors"
	"sort"
)

type IntervalMerger struct {
	intervals intervals
}

func NewIntervalMerger(intervals [][]int) (result *IntervalMerger, err error)  {
	result = &IntervalMerger{}
	for _, i := range intervals {
		i := interval(i)
		if !i.isValid() {
			return nil, errors.New(`intervals: Invalid interval: ` + IntervalToString(i))
		}
		result.intervals = append(result.intervals, i)
		sort.Sort(result.intervals)
	}
	return result, err
}

func(im *IntervalMerger) Merge() (err error) {
	var (
		i, j, k int
		result  intervals = make([]interval, 0)
	)

	for i < len(im.intervals) {
		k = i + 1
		for j = i + 1; j < len(im.intervals); j++ {
			if im.intervals[i].isMergeable(im.intervals[j]) {
				im.intervals[i].merge(im.intervals[j])
				k = j + 1
			}
		}
		result = append(result, im.intervals[i])
		i = k
	}
	im.intervals = result
	return nil
}

func(im *IntervalMerger) Intervals() (result[][]int) {
	for _, i := range im.intervals {
		result = append(result, []int(i))
	}
	return result
}
