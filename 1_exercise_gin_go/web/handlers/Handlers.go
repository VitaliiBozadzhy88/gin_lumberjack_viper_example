package handlers

import (
	"net/http"

	"exercises_go/1_exercise_gin_go/model"
	"github.com/gin-gonic/gin"
)

var cars = []model.Car{
	{"1", "toyota", "12000"},
	{"2", "suzuki", "9700"},
	{"3", "toyota", "18000"},
}

// GetCars show all cars
func GetCars(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, cars)
}

// PostCar adds new car in Cars
func PostCar(ctx *gin.Context) {
	var newCar model.Car

	if err := ctx.BindJSON(&newCar); err != nil {
		return
	}
	cars = append(cars, newCar)

	ctx.IndentedJSON(http.StatusCreated, newCar)
}

// GetCarById shows car by id
func GetCarById(ctx *gin.Context) {
	Id := ctx.Param("id")

	for _, car := range cars {
		if car.Id == Id {
			ctx.IndentedJSON(http.StatusOK, car)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}
