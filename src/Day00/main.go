package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sayAboutError(reason string) {
	fmt.Printf("ERROR: %s\n", reason)
}
func readValues(scanner *bufio.Scanner) []float64 {
	fmt.Println("input text, to stop adding, write \"\".")
	var arr []float64
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			sayAboutError("There was an error while translating your string to float," +
				"perhaps you wrote a number with *,* or the message was clearly not a number.\n" +
				"Number not added :c")
			continue
		}
		if value >= -100_000 && value <= 100_000 {
			arr = append(arr, value)
		} else {
			sayAboutError("Your number has been processed correctly," +
				"but it does not satisfy the interval specified in the task\n" +
				"Number not added :c")
		}
	}
	return arr
}
func readOutputMode(scanner *bufio.Scanner) int64 {
	fmt.Println("What information do you want to see?\n" +
		"0 -> All information\n" +
		"1 -> Mean value\n" +
		"2 -> Median value\n" +
		"3 -> Mode value\n" +
		"4 -> SD value")
	scanner.Scan()
	text := scanner.Text()
	value, err := strconv.ParseInt(text, 10, 64)
	if err != nil || !strings.Contains("01234", text) {
		sayAboutError("The number could not be converted or it is outside the suggested range, the value *0* is substituted.")
		value = 0
	}
	return value
}
func readFromConsole() ([]float64, int64) {
	scanner := bufio.NewScanner(os.Stdin)
	arr := readValues(scanner)
	mode := readOutputMode(scanner)

	if err := scanner.Err(); err != nil {
		sayAboutError("Scanner was unable to read your request or received non-standard input")
	}
	return arr, mode
}
func calcMean(arr []float64) float64 {
	var ans float64
	if len(arr) == 0 {
		ans = 0
	} else {
		for _, value := range arr {
			ans += value
		}
		ans = ans / float64(len(arr))
	}
	return ans
}
func calcMedian(arr []float64) float64 {
	var ans float64
	l := len(arr)
	if l == 0 {
		ans = 0
	} else if l%2 == 0 {
		ans = (arr[l/2-1] + arr[l/2]) / 2
	} else {
		ans = arr[l/2]
	}
	return ans
}
func calcMode(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	frequencyMap := make(map[float64]int)
	for _, value := range arr {
		frequencyMap[value]++
	}

	type kv struct {
		Key   float64
		Value int
	}
	var freqSlice []kv
	for key, value := range frequencyMap {
		freqSlice = append(freqSlice, kv{key, value})
	}

	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].Value > freqSlice[j].Value
	})

	maxValue := freqSlice[0].Value
	var maxValues []float64
	for _, kv := range freqSlice {
		if kv.Value == maxValue {
			maxValues = append(maxValues, kv.Key)
		} else {
			break
		}
	}
	sort.Float64s(maxValues)
	return maxValues[0]
}
func calcStandartDeviation(arr []float64, mean float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	var ans float64
	for _, value := range arr {
		ans += math.Pow(value-mean, 2)
	}
	ans = math.Sqrt(ans / float64(len(arr)))
	return ans
}
func majorStatisticalMetrics(arr []float64, mode int64) {
	var ansLine string
	sort.Float64s(arr)
	mean := calcMean(arr)

	if mode == 0 || mode == 1 {
		ansLine += fmt.Sprintf("Mean: %.2f\n", mean)
	}
	if mode == 0 || mode == 2 {
		ansLine += fmt.Sprintf("Median: %.2f\n", calcMedian(arr))
	}
	if mode == 0 || mode == 3 {
		ansLine += fmt.Sprintf("Mode: %.2f\n", calcMode(arr))
	}
	if mode == 0 || mode == 4 {
		ansLine += fmt.Sprintf("SD: %.2f\n", calcStandartDeviation(arr, mean))
	}
	fmt.Printf("%s", ansLine)
}

func main() {
	arr, value := readFromConsole()
	majorStatisticalMetrics(arr, value)
}
