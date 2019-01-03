package timeseries

import (
	"errors"
	floats "go-learning-codebin/numutils/floats"
	ints "go-learning-codebin/numutils/integers"
	"math"
)

func (series FloatSeries) getValues() []float64 {

	var values = make([]float64, 0, len(series))

	for _, value := range series {
		values = append(values, value)
	}

	return values

}

func (series FloatSeries) getTimes() []int64 {

	var times = make([]int64, 0, len(series))

	for time := range series {
		times = append(times, time)
	}

	return times

}

//Inefficient, we can store timeline in an ordered way with another implementation
func (series FloatSeries) firstAfter(time int64) (int64, error) {

	result, err := ints.Min(ints.Filter(series.getTimes(), func(x int64) bool {
		return x >= time
	}))

	if err != nil {
		return 0, errors.New("Cannot compute firstAfter because this series has not any coordinate in it")
	}

	return result, nil

}

//Inefficient, we can store timeline in an ordered way with another implementation
func (series FloatSeries) lastBefore(time int64) (int64, error) {

	result, err := ints.Max(ints.Filter(series.getTimes(), func(x int64) bool {
		return x <= time
	}))

	if err != nil {
		return 0, errors.New("Cannot compute lastBefore because this series has not any coordinate in it")
	}

	return result, nil

}

//ValueSum returns the sum of values inside the series
func (series FloatSeries) ValueSum() float64 {
	return floats.Sum(series.getValues())
}

//ValueAverage returns the average of values inside the series
func (series FloatSeries) ValueAverage() (float64, error) {

	res, err := floats.Avg(series.getValues())

	if err != nil {
		return 0, errors.New("Cannot compute Average, apparently this series hasn't any coordinate in it")
	}

	return res, nil

}

//ValueVariance returns the variance of the values contained inside the series
func (series FloatSeries) ValueVariance() (float64, error) {

	values := series.getValues()

	avg, err := floats.Avg(values)

	if err != nil {
		return 0, errors.New("Cannot compute Variance, apparently this series hasn't any coordinate in it")
	}

	varianceItems := floats.Map(values, func(x float64) float64 {
		return math.Pow(x-avg, 2)
	})

	return floats.Avg(varianceItems)

}

//ValueStdDeviation returns the standard deviation of the values contained inside the series
func (series FloatSeries) ValueStdDeviation() (float64, error) {

	variance, err := series.ValueVariance()

	if err != nil {
		return 0, errors.New("Cannot compute StdDeviation, apparently this series hasn't any coordinate in it")
	}

	return math.Sqrt(variance), nil

}

//ValueAt computes the value at a certain time as if this series emulates a "staircase" function
func (series FloatSeries) ValueAt(time int64) (float64, error) {

	key, err := series.lastBefore(time)
	altKey, altErr := series.firstAfter(time)

	if err != nil && altErr != nil {
		return 0, err
	} else if err != nil {
		return series[altKey], nil
	}

	return series[key], nil

}

//GetMin computes the minimum value of the entire series
func (series FloatSeries) GetMin() (float64, error) {
	return floats.Min(series.getValues())
}

//GetMax computes the maximum value of the entire series
func (series FloatSeries) GetMax() (float64, error) {
	return floats.Max(series.getValues())
}

//LerpAt computes the value at a certain time as if this series emulates a sawtooth-like function
func (series FloatSeries) LerpAt(time int64) (float64, error) {

	key, err := series.lastBefore(time)
	altKey, altErr := series.firstAfter(time)

	if err != nil && altErr != nil {
		return 0, err
	} else if err != nil {
		return series[altKey], nil
	} else if altErr != nil {
		return series[key], nil
	}

	dX := float64(altKey - key)
	dY := float64(series[altKey] - series[key])
	differential := dY / dX
	step := float64(time - key)

	return series[key] + step*differential, nil

}
