package main

import "fmt"

func main() {
	var friends1 = []string{"frizky", "sigit", "hassanudin", "inas salma", "iqbal hamdani", "kevin hugo", "medy kesuma putra", "mohammad fiqri", "ricky peater", "satrio utomo", "Septanus", "Clara Velita"}
	var friends2 = friends1[2:5]

	for i := range friends2 {
		fmt.Println(friends2[i])
	}
}
