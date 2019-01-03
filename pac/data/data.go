package pac

import (
	timeseries "go-learning-codebin/numutils/timeseries"
)

//History represents a temporal series of a PAC's value
type History struct {
	name   string
	series timeseries.FloatSeries
}

//NewHistory returns a new immutable instance of History
func NewHistory(name string, series timeseries.FloatSeries) History {
	return History{name, series}
}

//GetName gets the name of this history
func (history History) GetName() string {
	return history.name
}

//GetSeries gets the series of this history
func (history History) GetSeries() timeseries.FloatSeries {
	return history.series
}

//GetValueAt gives the value of the history at the given time
func (history History) GetValueAt(time int64) (float64, error) {
	return history.GetSeries().ValueAt(time)
}
