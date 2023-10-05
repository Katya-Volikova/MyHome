package main

import "fmt"

type Neighbours struct {
	name    string
	age     string
	course  string
	country string
	faculty string
}

func createNeighbour(name string, age string, course string, country string, faculty string) Neighbours {
	return Neighbours{
		name:    name,
		age:     age,
		course:  course,
		country: country,
		faculty: faculty,
	}
}

func initializeNeighbours() []Neighbours {
	var neighbours []Neighbours

	name := [3]string{"Katya", "Sasha", "Ada"}
	age := [3]string{"20", "21", "19"}
	course := [3]string{"3", "4", "1"}
	country := [3]string{"Russia", "Russia", "Nigeria"}
	faculty := [3]string{"IngeniaringAcademy", "Math", "Medicine"}
	neighboursMock := [][3]string{name, age, course, country, faculty}

	for i := 0; i < 3; i++ {
		neighbours = append(neighbours, createNeighbour(neighboursMock[0][i], neighboursMock[1][i], neighboursMock[2][i], neighboursMock[3][i], neighboursMock[4][i]))
	}

	return neighbours
}

func infoAboutNeighbours(neighbours []Neighbours) {
	fmt.Print("В комнате живет 3 девушки")
	for i := 0; i < 3; i++ {
		fmt.Print("\nИмя: ", neighbours[i].name, "\nЕй ", neighbours[i].age, "\nОна учится на ", neighbours[i].course, " курсе", " на факультете - ", neighbours[i].faculty, "\n")
	}
}
