package statistic

type StatisticInterface interface {
	Sum() float64
	Max() float64
	Min() float64
	Mean() float64
	Sort()
	GetData() []float64
	Median() float54
}
