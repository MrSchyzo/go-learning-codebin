package integers

import "errors"

//Map returns a new slice with every value mapped by the given function
func Map(ints []int64, f func(int64) int64) []int64 {

	var result = make([]int64, len(ints), cap(ints))

	for i := range ints {
		result[i] = f(ints[i])
	}

	return result

}

//Reduce reduces a slice into a single int64 value.
//Its second parameter is a function that accepts two values, the second value will be the i-th item of the slice
func Reduce(ints []int64, f func(int64, int64) int64, start int64) int64 {

	var result = start

	for i := range ints {
		result = f(result, ints[i])
	}

	return result

}

//Zip creates a merged slice of two int64 slices, merging them by providing a function as third parameter
func Zip(ints []int64, otherInts []int64, f func(int64, int64) int64) ([]int64, error) {

	if len(ints) != len(otherInts) {
		return nil, errors.New("Zip can only get two slices with the same length")
	}

	var result = make([]int64, len(ints))

	for i := range ints {
		result[i] = f(ints[i], otherInts[i])
	}

	return result, nil

}

//Sum computes the sum of a int64 slice
func Sum(ints []int64) int64 {

	return Reduce(ints, func(x, y int64) int64 {
		return x + y
	}, 0)

}

//Max computes the maximum of a int64 slice
func Max(ints []int64) (int64, error) {

	if len(ints) <= 0 {
		return 0, errors.New("Cannot compute Max because this slice hasn't any value in it")
	}

	result := Reduce(ints[1:], func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	}, ints[0])

	return result, nil

}

//Min computes the minimum of a int64 slice
func Min(ints []int64) (int64, error) {

	if len(ints) <= 0 {
		return 0, errors.New("Cannot compute Min because this slice hasn't any value in it")
	}

	result := Reduce(ints[1:], func(x, y int64) int64 {
		if x < y {
			return x
		}
		return y
	}, ints[0])

	return result, nil

}

//Avg computes the average of a int64 slice, it returns error if the given slice is empty
func Avg(ints []int64) (int64, error) {

	var length = len(ints)

	if length <= 0 {
		return 0, errors.New("Cannot compute an Avg of an empty slice")
	}

	return Sum(ints) / int64(length), nil

}

//Filter returns the filtered slice depending on the given condition
func Filter(ints []int64, condition func(int64) bool) []int64 {

	result := make([]int64, 0)

	for _, val := range ints {
		if condition(val) {
			result = append(result, val)
		}
	}

	return result

}
