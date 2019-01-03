package timeseries

//FloatCoordinate is just an immutable tuple (time: int64, value: float64)
type FloatCoordinate struct {
	time  int64 //time is how many milliseconds are elapsed after Epoch
	value float64
}

//NewFloatCoord creates a new FloatCoordinate immutable struct
func NewFloatCoord(time int64, value float64) FloatCoordinate {
	return FloatCoordinate{time, value}
}

//GetTime gets the time in milliseconds of a Float Coordinate
func (floatCoord FloatCoordinate) GetTime() int64 {
	return floatCoord.time
}

//GetValue gets the value of a Float Coordinate
func (floatCoord FloatCoordinate) GetValue() float64 {
	return floatCoord.value
}
