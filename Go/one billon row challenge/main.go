package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Measurements struct {
	Nome  string
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("meansurement.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer measurements.Close()
	scanner := bufio.NewScanner(measurements)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";")
		location := rawData[:semicolonIndex]
		temperature := rawData[semicolonIndex+1:]

		fmt.Println(location, temperature)
		return
	}

}
