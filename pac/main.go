package main

import (
	"fmt"
	"go-learning-codebin/numutils/floats"
	pac "go-learning-codebin/pac/data"
	"go-learning-codebin/pac/libs/internet"
	"os"
	"strconv"
	"time"
)

func getValueAtFromHistory(hist pac.History, givenTime time.Time) float64 {

	nowMillis := givenTime.Unix() * 1000

	value, err := hist.GetValueAt(nowMillis)

	if err != nil {
		panic(err.Error())
	}

	return value

}

func printValueAtFromHistory(hist pac.History, givenTime time.Time) {

	fmt.Printf(
		"%q has now (%q) value = %f\n",
		hist.GetName(),
		givenTime.Format(time.RFC3339),
		getValueAtFromHistory(hist, givenTime),
	)

}

func main() {

	url := "https://www.milanofinanza.it/Grafici/GraficoStorico?stockCodeRT=04fd&codiceStrumento=d23e"

	days := 0
	if len(os.Args) > 1 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			day = 0
		}
		days = day
	}

	hist, err := internet.HistoryFromURL(url)
	if err != nil {
		panic(err.Error())
	}

	start, err := time.Parse(time.RFC3339, "2017-01-16T09:00:00+00:00")
	if err != nil {
		panic(err.Error())
	}

	var values [24]float64
	for i := 0; i < 24; i++ {
		printValueAtFromHistory(hist, start.AddDate(0, i, 0))
		value := getValueAtFromHistory(hist, start.AddDate(0, i, 0))
		values[i] = 200.0 / value
	}

	totalExpenses := 200.0 * 24
	actualSingleValue := getValueAtFromHistory(hist, time.Now().AddDate(0, 0, days))
	averageSingleValue := totalExpenses / floats.Sum(values[:])
	actualValue := floats.Sum(values[:]) * actualSingleValue

	max, err := hist.GetSeries().GetMax()
	min, err := hist.GetSeries().GetMin()
	avg, err := hist.GetSeries().ValueAverage()
	variance, err := hist.GetSeries().ValueVariance()
	dev, err := hist.GetSeries().ValueStdDeviation()
	fmt.Println("Pac value:")
	fmt.Printf("Max = %f\n", max)
	fmt.Printf("Min = %f\n", min)
	fmt.Printf("Avg = %f\n", avg)
	fmt.Printf("StdDev = %f\n", dev)
	fmt.Printf("Var = %f\n", variance)
	fmt.Printf("Actual = %f\n", actualSingleValue)
	fmt.Printf("Avg owned = %f\n", averageSingleValue)
	fmt.Println("============================")
	fmt.Printf("Expenses = %f ; actual total value = %f ; gain = %f\n", totalExpenses, actualValue, actualValue-totalExpenses)

}
