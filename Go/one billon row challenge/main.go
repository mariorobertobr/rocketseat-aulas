package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	start := time.Now()
	measurements, err := os.Open("meansurement.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer measurements.Close()

	dados := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";")
		location := rawData[:semicolonIndex]
		rawtemp := rawData[semicolonIndex+1:]

		temperature, err := strconv.ParseFloat(rawtemp, 64)
		if err != nil {
			log.Fatal(err)
		}

		measurement, ok := dados[location]

		if !ok {
			measurement = Measurement{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temperature)
			measurement.Max = max(measurement.Max, temperature)
			measurement.Sum += temperature
			measurement.Count++

		}

		dados[location] = measurement

	}

	locations := make([]string, 0, len(dados))

	for name := range dados {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, name := range locations {
		measurement := dados[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f,", name, measurement.Min, measurement.Sum/float64(measurement.Count), measurement.Max)

	}

	fmt.Printf("}\n")

	// for name, measurement := range dados {
	// 	fmt.Printf("%s: %#+v\n", name, measurement)
	// }
	fmt.Println()
	fmt.Println(time.Since(start))

}
