package intervals

import (
	"strconv"
)

type intervals []interval

func (i intervals) Len() int {
	return len(i)
}

func (i intervals) Less( a, b int) bool {
	return i[a][0] < i[b][0]
}

func (i intervals) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func IntervalToString( i []int) string {
	result := `[`;
	for j, item := range i {
		result += strconv.Itoa(item)
		if j < len(i) - 1 {
			result += `,`
		}
	}
	result += `]`
	return result
}
