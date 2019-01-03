package json

import (
	json "encoding/json"
	"errors"
	"fmt"
	timeseries "go-learning-codebin/numutils/timeseries"
	pac "go-learning-codebin/pac/data"
	"strconv"
)

func generateHistoryFromStringyfiedCoordinates(coords [][]string) (map[int64]float64, error) {

	var history = make(map[int64]float64)

	for index, coord := range coords {

		time, timeErr := strconv.ParseInt(coord[0], 10, 64)
		value, valueErr := strconv.ParseFloat(coord[1], 64)

		if timeErr != nil {
			errorString := fmt.Sprintf("Cannot parse time of the coordinate at index %v", index+1)
			return nil, errors.New(errorString)
		} else if valueErr != nil {
			errorString := fmt.Sprintf("Cannot parse value of the coordinate at index %v", index+1)
			return nil, errors.New(errorString)
		}

		history[time] = value

	}

	return history, nil

}

//HistoryFromJSON returns a new History from the given JSON
func HistoryFromJSON(data []byte) (pac.History, error) {

	values := make([][]string, 0)

	err := json.Unmarshal(data, &values)

	if err != nil {
		return pac.History{}, errors.New("Cannot unmarshal given data")
	}

	return HistoryFromSlices(values)

}

//HistoryFromSlices returns a new History from the given 2D string
func HistoryFromSlices(values [][]string) (pac.History, error) {

	if len(values) == 0 {
		return pac.History{}, errors.New("Cannot unmarshal an empty 2D slice to a History")
	}

	var name = values[0][0]

	history, histErr := generateHistoryFromStringyfiedCoordinates(values[1:])

	if histErr != nil {
		return pac.History{}, histErr
	}

	return pac.NewHistory(name, timeseries.FloatSeriesFromMap(history)), nil

}
