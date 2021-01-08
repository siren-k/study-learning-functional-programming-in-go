package main

import (
	"bufio"
	"github.com/siren-k/study-learning-functional-programming-in-go/ch10/01_manad_workflow/workflow"
	"log"
	"os"
)

func main() {
	carCntr := 0
	if file, err := os.Open("/Users/benjamin/Lab/github/siren-k/study-learning-functional-programming-in-go/ch10/01_manad_workflow/cars.base64"); err == nil {
		defer file.Close()
		log.Println("----")
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			carCntr += 1
			log.Println("Processing car #", carCntr)
			line := scanner.Text()
			log.Println("IN :", line)
			err, carJson := workflow.ProcessCar(line)

			if err == nil {
				log.Println("OUT:", carJson)
			}

			log.Println("----")
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
