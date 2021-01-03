package main

import (
	"encoding/csv"
	"fmt"
	"github.com/wesovilabs/koazee"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CsvToStruct(filename string) []string {
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	csvFile, err := os.Open(fmt.Sprintf("%s/%s", pwd, filename))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var cars []string
	for _, car := range rawCSVdata {
		cars = append(cars, car[0])
	}
	return cars
}

func LoadCars() []string {
	return CsvToStruct("cars.csv")
}

func isDomestic(car string) bool {
	return strings.Contains(car, "Ford") || strings.Contains(car, "GM") || strings.Contains(car, "Chrysler")
}

// A continuum chain call is the execution of a series of functions
// so that the output of one function becomes the input of the next function.
func main() {
	cars := LoadCars()

	koazee.StreamOf(cars).
		Map(func(car string) string {
			fmt.Println("map ==>", car)
			return car
		}).
		Filter(func(car string) bool {
			match, _ := regexp.MatchString(".+[0-9].*", car)
			fmt.Println("filter1 ==> ", car, match)
			return match
		}).
		Filter(func(car string) bool {
			match := !isDomestic(car)
			fmt.Println("filter2 ==> ", car, match)
			return match
		}).
		ForEach(func(car string) {
			fmt.Println("forEach ==> ", car)
		}).Do()
}
