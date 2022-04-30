package statistic

import "sort"

type Statistic struct {
	Title string
	Datas []float64
}

/* Sum */
func (s *Statistic) Sum() float64 {
	var sum float64
	for _, data := range s.Datas {
		sum += data
	}
	return sum
}

/* Max */
func (s *Statistic) Max() float64 {
	var max float64
	for _, data := range s.Datas {
		if max < data {
			max = data
		}
	}
	return max
}

/* Min */
func (s *Statistic) Min() float64 {
	if len(s.Datas) == 0 {
		return 0
	}
	var min = s.Datas[0]
	for _, data := range s.Datas {
		if min > data {
			min = data
		}
	}
	return min
}

/* Mean */
func (s *Statistic) Mean() float64 {
	var count = float64(len(s.Datas))
	if count == 0 {
		return 0
	}
	var sum = s.Sum()
	mean := sum / count
	return mean
}

/* Sort */
func (s *Statistic) Sort() {
	sort.Sort(s)
}

/* Len */
func (s *Statistic) Len() int {
	return len(s.Datas)
}

/* Swap */
func (s *Statistic) Swap(i, j int) {
	s.Datas[i], s.Datas[j] = s.Datas[j], s.Datas[i]
}

/* Less */
func (s *Statistic) Less(i, j int) bool {
	return s.Datas[i] < s.Datas[j]
}
