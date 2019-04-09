package intervals

type interval []int

func (i *interval) isValid() bool {
	return len(*i) == 2
}

func (i *interval) isMergeable(interval interval) bool {
	if (*i)[0] <= interval[0] {
		return (*i)[1] >= interval[0] - 1
	}
	return (*i)[0] - 1 <= interval[1]
}

func (i *interval) merge(interval interval) {
	if (*i)[0] > interval[0] {
		(*i)[0] =  interval[0]
	}
	if (*i)[1] < interval[1] {
		(*i)[1] =  interval[1]
	}
}
