package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello World in Go")
	distance := 60.8 // distance in KMS
	fmt.Printf("The distance in miles %f \n", distance*0.621)
}
