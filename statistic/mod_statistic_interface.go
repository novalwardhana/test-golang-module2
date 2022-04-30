package statistic

type StatisticInterface interface {
	Sum() float64
	Max() float64
	Min() float64
	Mean() float64
}
