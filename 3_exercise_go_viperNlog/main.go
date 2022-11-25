package main

import (
	"exercises_go/3_exercise_go_viperNlog/repository"
)

func main() {
	repo := repository.NewRepository()

	//repo.Create("Suzuki", "4500")
	repo.Get()
	//fmt.Println(repository.Delete("17"))
}
