package timeseries

//FloatSeries is just a map (time: int64 -> value: float64)
type FloatSeries map[int64]float64

//FloatSeriesFromMap builds a FloatSeries from a int64 -> float64 map
func FloatSeriesFromMap(series map[int64]float64) FloatSeries {
	return series
}

//FloatSeriesFromCoordinates builds a FloatSeries from a slice of FloatCoordinate
func FloatSeriesFromCoordinates(coords []FloatCoordinate) FloatSeries {

	result := make(map[int64]float64)

	for _, coord := range coords {
		result[coord.time] = coord.value
	}

	return result

}

//GetCoordinates retrieves the coordinates slice of a FloatSeries
func (series FloatSeries) GetCoordinates() []FloatCoordinate {

	result := make([]FloatCoordinate, len(series))

	for key, value := range series {
		result = append(result, FloatCoordinate{key, value})
	}

	return result

}
