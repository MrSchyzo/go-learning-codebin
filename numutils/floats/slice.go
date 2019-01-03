package floats

import "errors"

//Map returns a new slice with every value mapped by the given function
func Map(floats []float64, f func(float64) float64) []float64 {

	var result = make([]float64, len(floats))

	for i := range floats {
		result[i] = f(floats[i])
	}

	return result

}

//Reduce reduces a slice into a single float64 value.
//Its second parameter is a function that accepts two values, the second value will be the i-th item of the slice
func Reduce(floats []float64, f func(float64, float64) float64, start float64) float64 {

	var result = start

	for i := range floats {
		result = f(result, floats[i])
	}

	return result

}

//Zip creates a merged slice of two float64 slices, merging them by providing a function as third parameter
func Zip(floats []float64, otherFloats []float64, f func(float64, float64) float64) ([]float64, error) {

	if len(floats) != len(otherFloats) {
		return nil, errors.New("Zip can only get two slices with the same length")
	}

	var result = make([]float64, len(floats))

	for i := range floats {
		result[i] = f(floats[i], otherFloats[i])
	}

	return result, nil

}

//Sum computes the sum of a float64 slice
func Sum(floats []float64) float64 {

	return Reduce(floats, func(x, y float64) float64 {
		return x + y
	}, 0)

}

//Max computes the maximum of a float64 slice
func Max(floats []float64) (float64, error) {

	if len(floats) <= 0 {
		return 0, errors.New("Cannot compute Max because this slice hasn't any value in it")
	}

	result := Reduce(floats[1:], func(x, y float64) float64 {
		if x > y {
			return x
		}
		return y
	}, floats[0])

	return result, nil

}

//Min computes the minimum of a float64 slice
func Min(floats []float64) (float64, error) {

	if len(floats) <= 0 {
		return 0, errors.New("Cannot compute Min because this slice hasn't any value in it")
	}

	result := Reduce(floats[1:], func(x, y float64) float64 {
		if x < y {
			return x
		}
		return y
	}, floats[0])

	return result, nil

}

//Avg computes the average of a float64 slice, it returns error if the given slice is empty
func Avg(floats []float64) (float64, error) {

	var length = len(floats)

	if length <= 0 {
		return 0, errors.New("Cannot compute an Avg of an empty slice")
	}

	return Sum(floats) / float64(length), nil

}

//Filter returns the filtered slice depending on the given condition
func Filter(floats []float64, condition func(float64) bool) []float64 {

	result := make([]float64, 0)

	for _, val := range floats {
		if condition(val) {
			result = append(result, val)
		}
	}

	return result

}
